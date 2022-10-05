package ggsignal

import (
	"github.com/guoyk93/gg"
	"os"
	"os/signal"
	"syscall"
)

// Guard guard the run stop func with signal handling, default handles SIGINT and SIGTERM
func Guard(run, stop func() error, sigs ...os.Signal) (err error) {
	if len(sigs) == 0 {
		sigs = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}

	chErr := make(chan error, 1)
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, sigs...)

	go func() {
		chErr <- run()
	}()

	select {
	case err = <-chErr:
		return
	case sig := <-chSig:
		gg.Log("signal caught: " + sig.String())
	}

	err = stop()
	return
}
