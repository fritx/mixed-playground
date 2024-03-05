package utils

import (
	"fmt"
	"strings"
	"testing"
)

// https://github.com/Jcowwell/go-algorithm-club/blob/go-main/Utils/assert.go
func preFormattedErrorString[T any](expected, got T) string {
	return fmt.Sprintf("expected: %+v got: %+v", expected, got)
}
func AssertEqual[T comparable](t *testing.T, value, expected T, msgs ...string) {
	// t.Run(fmt.Sprintf("AssertEqual - %v == %v", value, expected), func(t *testing.T) {
	if value != expected {
		s := preFormattedErrorString(expected, value)
		if len(msgs) > 0 {
			s = strings.Join(append(msgs, s), "\n")
		}
		t.Error(s)
	}
	// })
}
