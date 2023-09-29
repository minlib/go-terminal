package cmd

import (
	"fmt"
	"os/exec"
)

func Command(cmd string) *exec.Cmd {
	return exec.Command("cmd", "/c", cmd)
}

func Kill(pid int) string {
	return fmt.Sprintf("taskkill -f -pid %d", pid)
}
