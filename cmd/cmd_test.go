package cmd

import (
	"fmt"
	"os"
	"testing"
)

func TestCmd(t *testing.T) {
	s, err := Cmd("tasklist")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func TestKill(t *testing.T) {
	pid := os.Getpid()
	cmdStr := Kill(pid)
	s, err := Cmd(cmdStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
