package stress

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type stress struct {
	ctx context.Context
}

// New retursns new
func New(ctx context.Context) stress { // nolint: golint
	return stress{
		ctx: ctx,
	}
}

func (s stress) Handler(c echo.Context) error {
	nsecs := c.Param("secs")

	secs, err := strconv.ParseInt(nsecs, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nsecs)
	}

	go s.stresser(secs)

	return c.JSON(http.StatusOK, nsecs)
}

func (s stress) stresser(secs int64) {
	f, err := os.Open(os.DevNull)
	if err != nil {
		log.Error(err)
		return
	}
	defer f.Close()

	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)

	ctx, cancel := context.WithCancel(s.ctx)

	for i := 0; i < n; i++ {
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Fprintf(f, ".")
				}
			}
		}(ctx)
	}

	log.Infof("stressing for %d secs", secs)

	time.Sleep(time.Duration(secs) * time.Second)
	cancel()

	log.Info("releasing cpu's")
}
