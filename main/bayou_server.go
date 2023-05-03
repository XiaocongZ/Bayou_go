package main
import "Bayou/bayou"
//import "time"












func main() {


	//bayou.CallExample()
    s := bayou.MakeServer()
    go s.PeriodicCommits()
    go s.ListenKeystroke()
    //s.CallMessage(bayou.Operation{"main op"})

    s.Wg.Wait()

}
