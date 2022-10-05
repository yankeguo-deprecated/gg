package gg

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestMustEnv(t *testing.T) {
	require.Panics(t, func() {
		var out string
		MustEnv("HELLO", &out)
	})
	require.NotPanics(t, func() {
		out := "world"
		MustEnv("HELLO", &out)
		require.Equal(t, "world", out)
	})
	require.NotPanics(t, func() {
		var out string
		os.Setenv("HELLO", "world")
		MustEnv("HELLO", &out)
		require.Equal(t, "world", out)
	})
}
