package ggjson

import "encoding/json"

// Convert using json as intermedia format to convert one value to another
func Convert[T any](v any) (out T, err error) {
	var buf []byte
	if buf, err = json.Marshal(v); err != nil {
		return
	}
	if err = json.Unmarshal(buf, &out); err != nil {
		return
	}
	return
}
