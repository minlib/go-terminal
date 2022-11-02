package cmd

import (
	"fmt"
	"os/exec"

	"github.com/minlib/go-terminal/charset"
)

func WinCmd(cmd string) (string, error) {
	c := exec.Command("cmd", "/c", cmd)
	if buf, err := c.CombinedOutput(); err != nil {
		return "", err
	} else {
		output := charset.BytesToString(buf, charset.GB18030)
		return output, nil
	}
}

func WinKill(pid int) (string, error) {
	return WinCmd(fmt.Sprintf("taskkill -f -pid %d", pid))
}
