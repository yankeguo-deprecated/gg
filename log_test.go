package gg

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

type testLogger struct {
	sb *strings.Builder
}

func (t *testLogger) Log(v ...interface{}) {
	t.sb.WriteString(fmt.Sprint(v...))
}

func (t *testLogger) Logf(format string, v ...interface{}) {
	t.sb.WriteString(fmt.Sprintf(format, v...))
}

func TestLog(t *testing.T) {
	sb := &strings.Builder{}
	DefaultLogger = &testLogger{sb: sb}

	Log("Hello", "World\n")
	Logf("This is %d", 123)

	require.Equal(t, "HelloWorld\nThis is 123", sb.String())
}
