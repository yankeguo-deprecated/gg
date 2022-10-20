package ggsignal

import (
	"errors"
	"github.com/stretchr/testify/require"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestGuard(t *testing.T) {
	run := func() error {
		time.Sleep(time.Second)
		return errors.New("a")
	}
	stop := func() error {
		return nil
	}
	err := Guard(run, stop)
	require.Error(t, err)
	require.Equal(t, "a", err.Error())

	sig := make(chan struct{}, 1)

	run = func() error {
		t := time.After(time.Second * 10)
		select {
		case <-t:
		case <-sig:
		}
		return nil
	}
	stop = func() error {
		sig <- struct{}{}
		return nil
	}

	go func() {
		time.Sleep(time.Second)
		p, _ := os.FindProcess(os.Getpid())
		p.Signal(syscall.SIGINT)
	}()

	err = Guard(run, stop)
	require.NoError(t, err)
}
