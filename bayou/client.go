package bayou

import (
    "Bayou/tui"
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "time"
    "strconv"
    "golang.org/x/term"

)

const (
    lineLength = 5
)

type appState struct {
    self playerState
    oppo playerState
}

type playerState struct {
    pos int
    hp int
    fire bool
    fireat time.Time
}

func nextState(as appState, m Message, isSelf bool) appState {
    if isSelf {
        as.self = nextpState(as.self,m)
        if m.Op.Str == "f" && as.self.pos == as.oppo.pos {
            as.oppo.hp -= 1
            fmt.Println("\rshoot one")
            fmt.Println(m)
        }
    }else {
        as.oppo = nextpState(as.oppo,m)
        if m.Op.Str == "f" && as.self.pos == as.oppo.pos {
            as.self.hp -= 1
        }
    }
    return as
}

func nextpState(s playerState, m Message) playerState {
    switch m.Op.Str {
    case "a":
        s.pos -= 1
        if s.pos < 0 {
            s.pos = 0
        }
    case "d":
        s.pos += 1
        if s.pos >= lineLength {
            s.pos = lineLength - 1
        }
    case "f":
        s.fire = true
        s.fireat = m.Stp.AStp
    }
    //update s.fire
    if m.Stp.AStp.Sub(s.fireat)  > (time.Second * 3) {
        s.fire = false
    }
    return s
}

func (s *Server)ListenKeystroke() {

    state, err := term.MakeRaw(0)
    if err != nil {
        log.Fatalln("setting stdin to raw:", err)
    }
    defer func() {
        if err := term.Restore(0, state); err != nil {
            log.Println("warning, failed to restore terminal:", err)
        }
    }()

    in := bufio.NewReader(os.Stdin)
    for {
        r, _, err := in.ReadRune()
        if err != nil {
            log.Println("stdin:", err)
            break
        }
        fmt.Printf("read rune %q\r\n", r)
        switch r {
        case 'a':
            go s.CallMessage(Operation{"a"})
        case 'd':
            go s.CallMessage(Operation{"d"})
        case 'f':
            go s.CallMessage(Operation{"f"})
        }
        if r == 'q' {
            break
        }
    }
    s.Wg.Done()
}


func render(state appState){
    tui.Clear()
    view := "A Shoot Out in the Darkness\r\n"
    view += oppoLine(state.oppo.pos, state.oppo.fire) + "hp: " + strconv.Itoa(state.oppo.hp) + "\r\n"
    i := 5
    for i > 0 {
        view += ".. .. .. .. ..\r\n"
        i = i - 1
    }
    view += selfLine(state.self.pos) + "hp: " + strconv.Itoa(state.self.hp) + "\r\n"

    fmt.Println(view)


}

func selfLine(selfPos int) string {
    posSlice := strings.Split("## ## ## ## ##", " ")
    posSlice[selfPos] = "oo"
    return strings.Join(posSlice, " ")
}

func oppoLine(oppoPos int, justFired bool) string {
    if justFired {
        posSlice := strings.Split("## ## ## ## ##", " ")
        posSlice[oppoPos] = "oo"
        return strings.Join(posSlice, " ")
    }else {
        return "## ## ## ## ##"
    }
}
