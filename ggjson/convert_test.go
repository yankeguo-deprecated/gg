package ggjson

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTranslate(t *testing.T) {
	type T1 struct {
		A string `json:"a"`
	}

	type T2 struct {
		B string `json:"a"`
	}

	out, err := Convert[T2](T1{A: "bbb"})
	require.NoError(t, err)
	require.Equal(t, "bbb", out.B)
}
