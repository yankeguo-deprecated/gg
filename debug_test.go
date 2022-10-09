package gg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDebugSet_Debug(t *testing.T) {
	ds := NewDebugSet("")
	require.False(t, ds.Debug("aaa"))
	require.False(t, ds.Debug("bbb"))

	ds = NewDebugSet("echo:*,view:*,,ccc, ,")
	require.False(t, ds.Debug("aaa"))
	require.False(t, ds.Debug("bbb"))
	require.True(t, ds.Debug("echo:aa"))
	require.True(t, ds.Debug("echo"))
	require.True(t, ds.Debug("view:aa"))
	require.True(t, ds.Debug("view"))
	require.True(t, ds.Debug("ccc"))

	ds = NewDebugSet("*")
	require.True(t, ds.Debug("aaa"))
	require.True(t, ds.Debug("bbb"))
	require.True(t, ds.Debug("echo:aa"))
	require.True(t, ds.Debug("echo"))
	require.True(t, ds.Debug("view:aa"))
	require.True(t, ds.Debug("view"))
	require.True(t, ds.Debug("ccc"))

	ds = NewDebugSet("*,a,b,,,,c,   ,")
	require.True(t, ds.Debug("aaa"))
	require.True(t, ds.Debug("bbb"))
	require.True(t, ds.Debug("echo:aa"))
	require.True(t, ds.Debug("echo"))
	require.True(t, ds.Debug("view:aa"))
	require.True(t, ds.Debug("view"))
	require.True(t, ds.Debug("ccc"))
}
