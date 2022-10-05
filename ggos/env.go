package ggos

import (
	"errors"
	"os"
	"strings"
)

// MustEnv get environment variable from key, if both value and *out is empty, panic
func MustEnv(key string, out *string) {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		if *out == "" {
			panic(errors.New("missing environment variable '" + key + "'"))
		}
	} else {
		*out = val
	}
}
