package bash

import (
	"os/exec"
	"strings"
)

func Do(cmd string, args ...string) (string, error) {
	sh := exec.Command(cmd, args...)
	outBytes, err := sh.CombinedOutput()
	if err != nil {
		return "", err
	}
	out := strings.TrimSpace(string(outBytes))
	return out, err
}
