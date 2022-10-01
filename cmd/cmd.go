package cmd

import (
	"fmt"
	"os/exec"

	"github.com/minlib/go-terminal/charset"
)

func WinCmd(cmd string) (string, error) {
	c := exec.Command("cmd", "/c", cmd)
	outputByte, err := c.CombinedOutput()
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	}
	output := charset.BytesToString(outputByte, charset.GB18030)
	return output, nil
}
