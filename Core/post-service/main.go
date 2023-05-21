package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"post-service/internal/application"
	"post-service/internal/config"
	"post-service/logger"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("getting config failed: %s", err.Error())
	}

	optsLogger := logger.LoggerOptions{IsProd: cfg.IsProd}
	l, err := logger.New(optsLogger)
	if err != nil {
		log.Fatalf("logger initialization failed: %s", err.Error())
	}

	optsApp := application.AppOptions{DB_url: cfg.DbUrl, HTTP_port: cfg.HTTP_port, IsProd: cfg.IsProd, AuthURL: cfg.AuthURL}
	l.Info("DbUrl: " + cfg.DbUrl)
	l.Info("HTTP_port: " + fmt.Sprint(cfg.HTTP_port))
	l.Info("AuthURL: " + fmt.Sprint(cfg.AuthURL))
	app := application.New(optsApp)

	eg, groupCtx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		return app.Start()
	})

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case <-groupCtx.Done():
			l.Sugar().Error("app not started or a Server Error occurred: ", groupCtx.Err().Error())
			break loop
		}
	}

	stopCtx, stopCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer stopCancel()

	err = app.Stop(stopCtx)
	if err != nil {
		l.Sugar().Error(err)
	}
	l.Info("app stopped")
}
