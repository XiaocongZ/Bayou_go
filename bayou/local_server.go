package bayou
import "fmt"
import "time"
import "log"
import "net/rpc"
import "sync"

var stop bool = false
var stopMutex = sync.Mutex{}

type Server struct {
    sID int
    commits []Message
    accepts []Message
    acMutex sync.Mutex
    Wg sync.WaitGroup
    //aState appState
    cState appState
    aState appState
}
var initAppState appState = appState{playerState{0, 5, false, time.Time{}}, playerState{0, 5, false, time.Time{}}}
func MakeServer() *Server {
    s := Server{
        -1,
        make([]Message, 0),
        make([]Message, 0),
        sync.Mutex{},
        sync.WaitGroup{},
        initAppState,
        initAppState}
    s.Wg.Add(1)
    s.sID = CallReg()
    return &s
}

func (s *Server) PeriodicCommits(){
    for{
        s.acMutex.Lock()


        n := len(s.commits)
        newCommits := CallRecvCommits(n)
        s.commits = append(s.commits, newCommits...)

        //add to commits, delete from accepts

        for _, cm := range newCommits{
            s.cState = nextState(s.cState, cm, cm.Stp.SID == s.sID)
            //for now only same id will be in accepts
            if cm.Stp.SID == s.sID {
                //delete in c.aState
                for i, am := range s.accepts{
                    if cm.Stp.AStp != am.Stp.AStp || cm.Stp.SID != am.Stp.SID{
                        s.accepts = remove(s.accepts, i)
                        break
                    }
                }
            }
            if s.cState.self.hp <= 0 {
                stopMutex.Lock()
                stop = true
                stopMutex.Unlock()
                break
            }
            if s.cState.oppo.hp <= 0 {
                stopMutex.Lock()
                stop = true
                stopMutex.Unlock()
                break
            }


        }
        fmt.Println("newAccepts", s.accepts)
        s.aState = s.cState
        for _, m := range s.accepts{
            s.aState = nextState(s.aState, m, m.Stp.SID == s.sID)
        }
        fmt.Println("\rPeriodic")
        render(s.aState)
        s.acMutex.Unlock()

        time.Sleep(time.Second*3)
    }
}

func (s *Server) InstantAccepts(m Message){
    s.acMutex.Lock()
    defer s.acMutex.Unlock()
    s.accepts = append(s.accepts, m)
    s.aState = nextState(s.aState, m, m.Stp.SID == s.sID)
    fmt.Println("\rInstantAccept")
    render(s.aState)
}
//
// example function to show how to make an RPC call to the coordinator.
//
// the RPC argument and reply types are defined in rpc.go.
//
func CallExample() {

	// declare an argument structure.
	args := ExampleArgs{}

	// fill in the argument(s).
	args.X = 99

	// declare a reply structure.
	reply := ExampleReply{}

	// send the RPC request, wait for the reply.
	// the "Coordinator.Example" tells the
	// receiving server that we'd like to call
	// the Example() method of struct Coordinator.
	ok := call("Coordinator.Example", &args, &reply)
	if ok {
		// reply.Y should be 100.
		fmt.Println("reply.Y: ")
		fmt.Println(reply.Y)
	} else {
		fmt.Printf("call failed!\n")
	}
}


func CallReg() int {
    //reply := RegServerReply{}
    var reply int
    ok := call("Coordinator.RegServer", NilArgs{}, &reply)
    if ok {
		// reply.Y should be 100.
		return reply
	} else {
		fmt.Printf("call failed!\n")
        return -1
	}
}

func (s *Server) CallMessage(op Operation) {
    reply := Message{}
    stp := Stamp{time.Now(),time.Time{},s.sID}
    args := Message{stp,op}
    s.InstantAccepts(args)
    ok := call("Coordinator.Message", &args, &reply)
    if ok {
		// reply.Y should be 100.
		//fmt.Println("CallMessage: ")
		//fmt.Println("send", args, "\nrecv", reply)
	} else {
		fmt.Printf("call failed!\n")
	}
}

func CallRecvCommits(n int) []Message{
    //:= len(s.commits)
    reply := []Message{}
    ok := call("Coordinator.RecvCommits", &n, &reply)
    if ok {
        return reply
	} else {
		fmt.Printf("call failed!\n")
        return nil
	}
}

//
// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
//
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
