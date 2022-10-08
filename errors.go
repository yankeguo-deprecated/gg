package gg

import (
	"strconv"
	"strings"
	"sync"
)

type Errors []error

func (errs Errors) Error() string {
	sb := &strings.Builder{}
	for i, err := range errs {
		if err == nil {
			continue
		}
		if sb.Len() > 0 {
			sb.WriteString("; ")
		}
		sb.WriteRune('#')
		sb.WriteString(strconv.Itoa(i))
		sb.WriteString(": ")
		sb.WriteString(err.Error())
	}
	return sb.String()
}

type ErrorGroup struct {
	errors Errors
	locker *sync.RWMutex
}

func NewErrorGroup() *ErrorGroup {
	return &ErrorGroup{
		locker: &sync.RWMutex{},
	}
}

func (eg *ErrorGroup) Add(err error) {
	eg.locker.Lock()
	defer eg.locker.Unlock()

	eg.errors = append(eg.errors, err)
}

func (eg *ErrorGroup) Unwrap() error {
	eg.locker.RLock()
	defer eg.locker.RUnlock()

	for _, err := range eg.errors {
		if err != nil {
			return eg.errors
		}
	}

	return nil
}
