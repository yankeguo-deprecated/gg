package ggjson

import "encoding/json"

// Translate using json as intermedia format to translate one value to another
func Translate[T any](v any) (out T, err error) {
	var buf []byte
	if buf, err = json.Marshal(v); err != nil {
		return
	}
	if err = json.Unmarshal(buf, &out); err != nil {
		return
	}
	return
}
