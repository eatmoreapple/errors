package errors

import (
	"errors"
	"strings"
)

type wrapper struct {
	err error
	msg string
}

// Error implements the error interface.
func (w *wrapper) Error() string {
	return w.msg
}

// Unwrap returns the wrapped error.
func (w *wrapper) Unwrap() error {
	return w.err
}

// Warp wraps the given error with the given message.
func Warp(err error, msg string) error {
	return &wrapper{err: err, msg: msg}
}

// Contains returns true if the error group contains the given error.
func Contains(err error, msg string) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), msg)
}

// ErrorGroup defines a group of errors.
type ErrorGroup []error

// Contains returns true if the error group contains the given error.
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

// Uniq returns a new ErrorGroup with unique errors.
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

// Len returns the number of errors in the group.
func (eg ErrorGroup) Len() int {
	return len(eg)
}

// WarpALL wraps all errors in the group with the given message.
func (eg ErrorGroup) WarpALL(err error) ErrorGroup {
	if len(eg) == 0 {
		return nil
	}
	var group ErrorGroup
	for _, err2 := range eg {
		group = append(group, Warp(err2, err.Error()))
	}
	return group
}
