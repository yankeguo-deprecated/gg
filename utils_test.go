package gg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepeat(t *testing.T) {
	require.Equal(t, []string{"a", "a"}, Repeat(2, "a"))
	require.Equal(t, []int{5, 5}, Repeat(2, 5))
}

func TestPtr(t *testing.T) {
	v := 2
	nv := Ptr(v)
	v = 3
	require.Equal(t, 2, *nv)
}
