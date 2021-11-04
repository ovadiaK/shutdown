package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ovadiaK/shutdown/internal/input"
	"github.com/ovadiaK/shutdown/internal/server"
)

func main() {

	s := server.New(
		input.New(),
	)

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	go func() {
		defer func() {
			stop()
			cancel()
		}()

		fmt.Println("waiting for signal...")

		<-ctx.Done()

		fmt.Println("signal received")
	}()

	s.Run(ctx)
}
