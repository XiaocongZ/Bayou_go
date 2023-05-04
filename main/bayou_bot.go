package main
import "Bayou/bayou"
import "time"
import "math/rand"












func main() {


	//bayou.CallExample()
    s := bayou.MakeServer()
    actions := []string{"a", "d", "f"}
    for {
        action := actions[rand.Intn(3)]
        s.CallMessage(bayou.Operation{action})

        time.Sleep(time.Second * 1)
    }



}
