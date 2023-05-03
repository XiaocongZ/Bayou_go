package tui

import (
    "os/exec"
    "os"
)

func Clear(){
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}
