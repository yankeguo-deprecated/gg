package ggjson

import (
	"encoding/json"
)

func Unmarshal[T any](buf []byte) (out T, err error) {
	err = json.Unmarshal(buf, &out)
	return
}

func MarshalPretty(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
