package ggos

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestHostnameSequenceID(t *testing.T) {
	err := os.Setenv("HOSTNAME", "aaa-12")
	require.NoError(t, err)

	id, err := HostnameSequenceID()
	require.NoError(t, err)
	require.Equal(t, 12, id)

	err = os.Setenv("HOSTNAME", "aaa.11")
	require.NoError(t, err)

	id, err = HostnameSequenceID()
	require.NoError(t, err)
	require.Equal(t, 11, id)

	err = os.Setenv("HOSTNAME", "aaa_13")
	require.NoError(t, err)

	id, err = HostnameSequenceID()
	require.NoError(t, err)
	require.Equal(t, 13, id)
}
