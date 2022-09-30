package gg

import (
	"context"
	"fmt"
)

// Guard recover from panic and set err
func Guard(err *error) {
	if r := recover(); r != nil {
		if re, ok := r.(error); ok {
			*err = re
		} else {
			*err = fmt.Errorf("panic: %v", r)
		}
	}
}

// MustContext panic ctx.Err() if not nil
func MustContext(ctx context.Context) {
	if ctx.Err() != nil {
		panic(ctx.Err())
	}
}

// Must0 panic err if not nil
func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

// Must panic err if not nil, return remaining values
func Must[T any](v T, err error) T {
	if err == nil {
		return v
	} else {
		panic(err)
	}
}
