package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/boletia/rest-cpu-load/pkg/stress"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	wg.Add(1)
	e := echo.New()
	go start(ctx, e)
	go stop(ctx, cancel, e, &wg)

	wg.Wait()
}

func start(ctx context.Context, e *echo.Echo) {
	service := stress.New(ctx)

	e.GET("/stress/:secs", service.Handler)
	e.Start(":8080")
}

func stop(ctx context.Context, cancel context.CancelFunc, e *echo.Echo, wg *sync.WaitGroup) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case <-interrupt:
		echoCtx, cancel := context.WithCancel(ctx)
		e.Shutdown(echoCtx)

		cancel()
		wg.Done()
	}
}
