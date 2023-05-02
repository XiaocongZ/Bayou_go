package main


import "Bayou/bayou"
import (
    "os/exec"
    "os"
)
import "fmt"
import "time"


func main() {


	b := bayou.MakeCoordinator()


    fmt.Println(b)
    n := 0
    for  {
        fmt.Println("waiting", n)
        time.Sleep(time.Second*2)
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
        n = n + 1
    }
}
