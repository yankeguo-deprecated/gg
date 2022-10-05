package ggsignal

import (
	"github.com/guoyk93/gg"
	"os"
	"os/signal"
	"syscall"
)

func Guard(run, stop func() error) (err error) {
	chErr := make(chan error, 1)
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		chErr <- run()
	}()

	select {
	case err = <-chErr:
		return
	case sig := <-chSig:
		gg.Log("signal caught:", sig.String())
	}

	err = stop()
	return
}
