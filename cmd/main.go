package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/morris-zheng/go-slim/internal/conf"
	"github.com/morris-zheng/go-slim/internal/delivery"
	"github.com/morris-zheng/go-slim/internal/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	// load config
	c := conf.Load("config.yaml")

	// gin router
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// new service context
	svc := domain.NewServiceContext(c)
	// register delivery
	delivery.Register(svc, r)

	// server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", c.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
