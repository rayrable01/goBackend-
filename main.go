package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	sigchan := make(chan os.Signal, 1)

	signal.Notify(sigchan, syscall.SIGTERM)

	go func() {
		<-sigchan
		// При отмене контекста происходит gracefull shutdown
		cancel()
	}()
}
