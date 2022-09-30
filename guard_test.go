package gg

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGuard(t *testing.T) {
	var err error
	func() {
		defer Guard(&err)
		panic(errors.New("hello"))
	}()
	require.Error(t, err)
	require.Equal(t, "hello", err.Error())
}

func TestMustContext(t *testing.T) {
	var err error
	func() {
		defer Guard(&err)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		MustContext(ctx)
	}()
	require.Error(t, err)
	require.Equal(t, context.Canceled, err)
}

func TestMust0(t *testing.T) {
	var err error
	func() {
		defer Guard(&err)
		Must0(errors.New("hello"))
	}()
	require.Error(t, err)
	require.Equal(t, "hello", err.Error())
}

func TestMust(t *testing.T) {
	{
		var err error
		var (
			v1 int
		)
		func() {
			defer Guard(&err)
			v1 = Must(3, errors.New("hello"))
		}()
		require.Error(t, err)
		require.Equal(t, "hello", err.Error())
		require.Equal(t, 0, v1)
	}
	{
		var err error
		var (
			v1 int
		)
		func() {
			defer Guard(&err)
			v1 = Must(3, nil)
		}()
		require.NoError(t, err)
		require.Equal(t, 3, v1)
	}
}
