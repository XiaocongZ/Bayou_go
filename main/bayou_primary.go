package main


import "Bayou/bayou"
import "Bayou/tui"

import "fmt"
import "time"


func main() {


	b := bayou.MakeCoordinator()

    n := 0
    for  {
        fmt.Println("waiting", n)
		b.PrintCommits()
        time.Sleep(time.Second*5)
        tui.Clear()
        n = n + 1
    }
}
