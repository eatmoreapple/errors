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

func TestUniq(t *testing.T) {
	var err1 = errors.New("error1")
	var err2 = errors.New("error2")
	var eg = ErrorGroup{err1, err2, err1, err2, err2}.Uniq()
	if len(eg) != 2 {
		t.Error("un uniqed")
	}
}
