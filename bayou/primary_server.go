package bayou

import "log"
import "net"
import "os"
import "net/rpc"
import "net/http"
import "fmt"
import "time"
import "sync"

type Coordinator struct {
    sID int
    smutex sync.Mutex //mutex for writing sID
    messageChan chan Message // buffered channel for Message, as a concurrent FIFO queue
    commits []Message
}

func (c *Coordinator) PrintCommits(){
    fmt.Println(len(c.commits), c.commits)
}

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	//reply.Y = args.X + 1
	reply.Y = []int{1, 2, 3}
	return nil
}


func (c *Coordinator) RegServer(args *ExampleArgs, reply *int) error {

    c.smutex.Lock()
	*reply = c.sID
    c.sID = c.sID + 1
    c.smutex.Unlock()
	return nil
}

func (c *Coordinator) Message(args *Message, reply *Message) error {
    time.Sleep(time.Second*2)
    //(*args).Stp.CStp = time.Now()
    c.messageChan <- *args
    fmt.Println(*args)
    *reply = *args
    return nil
}

//no concurrency control needed
func (c *Coordinator) RecvCommits(n int, reply *[]Message) error {
    *reply = c.commits[n:]
    return nil
}

func (c *Coordinator) handleMessage(){
    for {
        m := <- c.messageChan
        m.Stp.CStp = time.Now()
        c.commits = append(c.commits, m)
    }
}


//
// start a thread that listens for RPCs from worker.go
//
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeCoordinator() *Coordinator {
	c := Coordinator{
        1,
        sync.Mutex{},
        make(chan Message, 200),
        make([]Message,0)}

	// Your code here.

	//start multi-threading
	c.server()
    go c.handleMessage()
	return &c
}
