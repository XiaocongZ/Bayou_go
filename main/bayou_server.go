package main
import "Bayou/bayou"
import "time"










func main() {

	//bayou.CallExample()
    s := bayou.MakeServer()
    go s.PeriodicCommits()
    i := 10
    for i > 0 {
        bayou.CallMessage()
        time.Sleep(time.Second*1)
        }
    time.Sleep(time.Second*10)
}
