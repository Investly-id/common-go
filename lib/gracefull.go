package lib

import (
	"os"
	"os/signal"
	"syscall"
)

type GracefullShutdown struct {
	channel chan os.Signal
}

func NewGracefullShutdown() *GracefullShutdown {
	chanServer := make(chan os.Signal, 1)
	signal.Notify(chanServer, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	return &GracefullShutdown{
		channel: chanServer,
	}
}

func (sc *GracefullShutdown) Wait() {
	defer close(sc.channel)

	<-sc.channel
	signal.Stop(sc.channel)
}
