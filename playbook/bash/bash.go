package bash

import (
	"os/exec"
)

func Do(cmd string, args ...string) bool {
	sh := exec.Command(cmd, args...)
	_, err := sh.Output()

	if werr, ok := err.(*exec.ExitError); ok {
		if s := werr.Error(); s != "0" {
			return false
		}
	}
	return true
}
