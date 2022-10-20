package gg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPtr(t *testing.T) {
	v := 2
	nv := Ptr(v)
	v = 3
	require.Equal(t, 2, *nv)
}

func TestEq(t *testing.T) {
	require.True(t, Eq(1)(1))
	require.False(t, Eq(1)(2))
}

func TestNeq(t *testing.T) {
	require.False(t, Neq(1)(1))
	require.True(t, Neq(1)(2))
}
