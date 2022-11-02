package cmd

import (
	"fmt"
	"os"
	"testing"
)

func TestWinCmd(t *testing.T) {
	s, err := WinCmd("tasklist")
	// s, err := WinCmd("ipconfig")
	// s, err := WinCmd("dir")
	fmt.Println(s, err)
}

func TestWinKill(t *testing.T) {
	pid := os.Getpid()
	s, err := WinKill(pid)
	fmt.Println("pid:", pid)
	fmt.Println(s, err)
}
