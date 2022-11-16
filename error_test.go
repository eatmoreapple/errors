package errors

import (
	"errors"
	"testing"
)

func TestWarp(t *testing.T) {
	var target = errors.New("target error")
	err := Warp(target, "my error")
	if !errors.Is(err, target) {
		t.Fatal("not equal")
	}
}
