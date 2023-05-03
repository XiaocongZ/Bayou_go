package bayou

//
// RPC definitions.
//
// remember to capitalize all names.
//

import "os"
import "strconv"
import "time"

func remove[T any](slice []T, s int) []T {
    return append(slice[:s], slice[s+1:]...)
}

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y []int
}

type NilArgs struct {}

type RequestTaskReply struct {
	Exit bool
	IsMapTask bool
	IsReduceTask bool
	Files []string
	ReduceN int //reduce number or partition number
}

type FinishArgs RequestTaskReply

// Add your RPC definitions here.
type Stamp struct {
    AStp time.Time
    CStp time.Time
    SID int
}

type Operation struct {
	Str string
}

type Message struct {
    Stp Stamp
    Op Operation
}

type RegServerReply struct {
    SID int
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/Distributed-Algorithm-Bayou-"
	s += strconv.Itoa(os.Getuid())
	return s
}
