package bayou
import "fmt"
import "time"
import "log"
import "net/rpc"


type Server struct {
    commits []Message
}

func MakeServer() *Server {
    s := Server{
        make([]Message, 0)}
    return &s
}

func (s *Server) PeriodicCommits(){
    for{
        n := len(s.commits)
        newCommits := CallRecvCommits(n)
        s.commits = append(s.commits, newCommits...)
        fmt.Println(len(s.commits), s.commits)
        time.Sleep(time.Second*2)
    }
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


func CallReg() {
    reply := RegServerReply{}
    ok := call("Coordinator.RegServer", NilArgs{}, &reply)
    if ok {
		// reply.Y should be 100.
		fmt.Println("CallReg: ")
		fmt.Println(reply)
	} else {
		fmt.Printf("call failed!\n")
	}
}

func CallMessage() {
    reply := Message{}
    stp := Stamp{time.Now(),time.Time{},0}
    args := Message{stp,Operation{}}
    ok := call("Coordinator.Message", &args, &reply)
    if ok {
		// reply.Y should be 100.
		fmt.Println("CallMessage: ")
		fmt.Println("send", args, "\nrecv", reply)
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
