package helper

import (
	"testing"
)

func TestFindVariable(t *testing.T) {
	val := "{{var}}"
	_, res := FindVariable(val)

	if "var" != res {
		t.Errorf("expecting %s but found %s", "var", res)
	}
}
