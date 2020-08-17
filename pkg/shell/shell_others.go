// +build !linux,!darwin,!windows

package shell

import (
	"errors"
	"os/exec"
)

var ErrNotImplemented = errors.New("shell: not implemented")

func GetCmd(s string) (*exec.Cmd, error) {
	return nil, ErrNotImplemented
}