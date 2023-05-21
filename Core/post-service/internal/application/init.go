package application

import (
	"context"
	"fmt"
	"post-service/internal/adapters/db"
	"post-service/internal/adapters/http"
	"post-service/internal/domain/usecases"

	"go.uber.org/zap"
)

type App struct {
	opts          AppOptions
	shutdownFuncs []func(ctx context.Context) error
}

type AppOptions struct {
	DB_url    string
	HTTP_port int
	IsProd    bool
	AuthURL   string
}

func New(opts AppOptions) *App {
	return &App{
		opts: opts,
	}
}

func (app *App) Start() error {

	log := zap.L()
	log.Info("application started")

	userStorage, err := db.New(context.Background(), app.opts.DB_url)
	if err != nil {
		log.Info("create user storage failed: " + err.Error())
		return fmt.Errorf("create post storage failed: %w", err)
	}
	log.Info("user storage created")

	post_service := usecases.New(userStorage)

	optsAdapter := http.AdapterOptions{HTTP_port: app.opts.HTTP_port, IsProd: app.opts.IsProd, AuthURL: app.opts.AuthURL}
	s, err := http.New(post_service, optsAdapter)
	if err != nil {
		return fmt.Errorf("server not started %w", err)
	}

	app.shutdownFuncs = append(app.shutdownFuncs, s.Stop)
	err = s.Start()
	if err != nil {
		return fmt.Errorf("server not started: %w", err)
	}
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	var err error
	for i := len(a.shutdownFuncs) - 1; i >= 0; i-- {
		err = a.shutdownFuncs[i](ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
