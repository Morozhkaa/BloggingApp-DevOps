package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-app/internal/application"
	"user-app/internal/config"
	"user-app/pkg/infra/logger"
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

	optsApp := application.AppOptions{DB_url: cfg.DbUrl, HTTP_port: cfg.HTTP_port, IsProd: cfg.IsProd}
	app := application.New(optsApp)
	err = app.Start()
	if err != nil {
		l.Sugar().Fatalf("app not started: %s", err.Error())
	}

	<-ctx.Done()

	stopCtx, stopCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer stopCancel()

	err = app.Stop(stopCtx)
	if err != nil {
		l.Sugar().Error(err)
	}
	l.Info("app stopped")
}
