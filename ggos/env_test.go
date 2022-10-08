package ggos

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

func TestBoolEnv(t *testing.T) {
	var out bool
	os.Setenv("TEST_A", "true")
	BoolEnv("TEST_A", &out)
	require.True(t, out)
	BoolEnv("TEST_B", &out)
	require.True(t, out)
}
