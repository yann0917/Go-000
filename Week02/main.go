package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/yann0917/Go-000/Week02/internal/dao"
	"github.com/yann0917/Go-000/Week02/router"
)

func main() {
	defer dao.DB.Close()

	r := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8000),
		Handler:        r,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := r.Run(":" + "8000"); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown error:", err)
	}

	log.Println("Server exiting")
}
