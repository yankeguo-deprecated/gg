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
