package cmd

import (
	"os/exec"
	"strconv"
)

func Cmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	return cmd.Run()
}

func Kill(pid int) error {
	return Cmd("kill", "-1", strconv.Itoa(pid))
}
