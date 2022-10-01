package cmd

import (
	"fmt"
	"testing"
)

func TestWinCmd(t *testing.T) {
	s, err := WinCmd("tasklist")
	fmt.Println(s, err)
	//WinCmd("ipconfig")
	//WinCmd("dir")
}
