package signal

import (
	"os"
	"os/signal"
	"syscall"
)

var exitSignals = []os.Signal{os.Interrupt, syscall.SIGTERM} // SIGTERM is POSIX specific

// RegisterExitHandlers returns a stop channel to wait on exit signal
func RegisterExitHandlers(shutdownFuncs ...func()) (stop chan struct{}) {
	stop = make(chan struct{})
	s := make(chan os.Signal, len(exitSignals))
	signal.Notify(s, exitSignals...)

	go func() {
		// Wait for a signal from the OS before dispatching
		// a stop signal to all other goroutines observing this channel.
		<-s
		close(stop)

		// execute our shutdown functions
		for _, f := range shutdownFuncs {
			f()
		}
	}()

	return stop
}

/* Sample Usage:
    _, cancel := context.WithCancel(context.Background())
    defer cancel()
    stop := signal.RegisterExitHandlers(cancel)
...
	<-stop
	cancel()
	log.Info().Msg("Stopping...")
*/
