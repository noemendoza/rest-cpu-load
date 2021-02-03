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
	var wg sync.WaitGroup

	ctx := context.Background()

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	wg.Add(1)

	e := echo.New()

	go start(ctx, e)

	go stop(ctx, e, &wg)

	wg.Wait()
}

func start(ctx context.Context, e *echo.Echo) {
	service := stress.New(ctx)

	e.GET("/stress/:secs", service.Handler)

	if err := e.Start(":8080"); err != nil {
		log.Errorf("start server error %s", err)
	}
}

func stop(ctx context.Context, e *echo.Echo, wg *sync.WaitGroup) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	<-interrupt

	echoCtx, cancel := context.WithCancel(ctx)
	if err := e.Shutdown(echoCtx); err != nil {
		log.Errorf("shutdown server error %s", err)
	}

	cancel()
	wg.Done()
}
