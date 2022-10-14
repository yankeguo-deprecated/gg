package ggjson

import (
	"github.com/guoyk93/gg"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshal(t *testing.T) {
	type valType struct {
		Hello string `json:"hello"`
	}
	buf := []byte(`{"hello":"world"}`)
	val, err := Unmarshal[valType](buf)
	require.NoError(t, err)
	require.Equal(t, "world", val.Hello)
}

func TestMarshalPretty(t *testing.T) {
	v := gg.Must(MarshalPretty(map[string]interface{}{"hello": "world"}))
	require.Equal(t, `{
  "hello": "world"
}`, string(v))
}
