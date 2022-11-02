package cmd

import (
	"fmt"
	"os"
	"testing"
)

func TestCmd(t *testing.T) {

}

func TestKill(t *testing.T) {
	pid := os.Getpid()
	err := Kill(pid)
	fmt.Println("pid:", pid)
	fmt.Println(err)
}
