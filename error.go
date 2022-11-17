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

func (eg ErrorGroup) Uniq() ErrorGroup {
	if len(eg) == 0 {
		return nil
	}
	if len(eg) == 1 {
		return eg
	}
	var group ErrorGroup
	for _, err := range eg {
		var found bool
		for _, err2 := range group {
			if err == err2 {
				found = true
				break
			}
		}
		if !found {
			group = append(group, err)
		}
	}
	return group
}

func (eg ErrorGroup) Len() int {
	return len(eg)
}
