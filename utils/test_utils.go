package utils

import (
	"fmt"
	"testing"
)

func Fail(t *testing.T, format string, args ...interface{}) {
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("Failed: "+format, args...))
	t.Fail()
}

func Pass(format string, args ...interface{}) {
	fmt.Printf("\x1b[32;1m%s\x1b[0m\n", fmt.Sprintf("Passed: "+format, args...))
}
