package errors

import (
	"errors"
	"strings"
)

type wrapper struct {
	err error
	msg string
}

func (w *wrapper) Error() string {
	return w.msg
}

func (w *wrapper) Unwrap() error {
	return w.err
}

func Warp(err error, msg string) error {
	return &wrapper{err: err, msg: msg}
}

func Contains(err error, msg string) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), msg)
}

type ErrorGroup []error

func (eg ErrorGroup) Contains(target error) bool {
	if len(eg) == 0 {
		return false
	}
	for _, err := range eg {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}
