package cmd

import (
	"fmt"
	"os/exec"
)

func Command(cmd string) *exec.Cmd {
	return exec.Command("/bin/bash", "-c", cmd)
}

func Kill(pid int) string {
	return fmt.Sprintf("kill -9 %d", pid)
}
