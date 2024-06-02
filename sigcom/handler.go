package sigcom

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func handle() {
	ctx := context.Background()
	sigctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	toctx, cancel2 := context.WithTimeout(sigctx, 5*time.Second)
	defer cancel2()

	select {
	case <-sigctx.Done():
		fmt.Println("signal received")
	case <-toctx.Done():
		fmt.Println("timeout")
	}
}
