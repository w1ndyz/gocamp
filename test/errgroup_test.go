package test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"

	h "github.com/w1ndyz/gocamp/ch03/http"
	"golang.org/x/sync/errgroup"
)

// Test .
func TestErrGroup(t *testing.T) {
	var done = make(chan int)
	group, ctx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/stop", func(rw http.ResponseWriter, r *http.Request) {
			done <- 1
		})
		s := h.NewServer(":3000", mux)
		go func() {
			err := s.Start()
			if err != nil {
				fmt.Println(err)
			}
		}()

		select {
		case <-done:
			return s.Stop()
		case <-ctx.Done():
			return fmt.Errorf("all close")
		}
	})

	group.Go(func() error {
		signals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
		sigCh := make(chan os.Signal, len(signals))
		signal.Notify(sigCh, signals...)
		for {
			fmt.Println("signal...")
			select {
			case <-ctx.Done():
				fmt.Println("all close")
				return ctx.Err()
			case <-sigCh:
				return fmt.Errorf("signal close")
			}
		}
	})

	fmt.Println("start to catch err.....")
	err := group.Wait()
	fmt.Println("err is:", err)
}
