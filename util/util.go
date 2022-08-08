package util

import (
	"errors"
	"os/exec"
)

func DownloadFile() error {
	return nil
}

func ExecCommand(cmdName string, args []string) error {
	cmd := exec.Command(cmdName)
	err := cmd.Run()
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
