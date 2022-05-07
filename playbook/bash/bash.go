package bash

import (
	"github.com/fatih/color"
	"os/exec"
)

func prepareToRun(path string) bool {
	return Do("chmod", "+x", path)
}

func Run(path string) bool {
	if !prepareToRun(path) {
		return false
	}
	sh := exec.Command("/bin/sh", path)
	_, err := sh.Output()

	if werr, ok := err.(*exec.ExitError); ok {
		if s := werr.Error(); s != "0" {
			color.Red("Script run error", err)
			return false
		}
	}

	return true
}

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
