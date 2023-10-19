package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/arieffian/simple-commerces-monorepo/internal/config"
	"github.com/arieffian/simple-commerces-monorepo/internal/servers"
)

const (
	shutDownTimeout = 10 * time.Second
	appTimeout      = 30 * time.Second
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config. %+v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), appTimeout)
	defer cancel()

	log.Infof("Starting %s server...", cfg.Service)

	server, err := servers.NewServer(ctx, *cfg)
	if err != nil {
		log.Fatalf("failed to create the new server: %s\n", err)
	}

	go func() {
		if err := server.Listen(cfg.APIAddress); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancelShutdown := context.WithTimeout(context.Background(), shutDownTimeout)
	defer cancelShutdown()

	<-c
	fmt.Println("Gracefully shutting down...")
	_ = server.Shutdown(ctx)

	fmt.Println("Fiber was successful shutdown.")
}
