package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	group := new(errgroup.Group)
	serverErr := make(chan error, 1)
	quit := make(chan os.Signal)

	s := http.Server{
		Addr:           ":8099",
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	group.Go(func() error {
		go func() {
			serverErr <- s.ListenAndServe()
		}()
		select {
		case err := <-serverErr:
			close(quit)
			close(serverErr)
			return err
		}
	})

	group.Go(func() error {
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			log.Fatal("server shutdown", err)
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Printf("get err: %v", err.Error())
	}
}