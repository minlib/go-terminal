package cmd

import (
	"github.com/minlib/go-terminal/charset"
	"os/exec"
)

func Cmd(cmd string) (string, error) {
	c := Command(cmd)
	return Output(c)
}

func Output(c *exec.Cmd) (string, error) {
	bytes, err := c.CombinedOutput()
	if err != nil {
		return "", err
	}
	output := charset.BytesToString(bytes, charset.GB18030)
	return output, nil
}
