package application

import (
	"context"
	"fmt"
	"user-app/internal/adapters/db"
	"user-app/internal/adapters/http"
	"user-app/internal/domain/usecases"
)

type App struct {
	opts          AppOptions
	shutdownFuncs []func(ctx context.Context) error
}

type AppOptions struct {
	DB_url    string
	HTTP_port int
	IsProd    bool
}

func New(opts AppOptions) *App {
	return &App{
		opts: opts,
	}
}

func (app *App) Start() error {

	userStorage, err := db.New(context.Background(), app.opts.DB_url)
	if err != nil {
		return fmt.Errorf("create user storage failed: %w", err)
	}

	auth := usecases.New(userStorage)

	optsAdapter := http.AdapterOptions{HTTP_port: app.opts.HTTP_port, IsProd: app.opts.IsProd}
	s, err := http.New(auth, optsAdapter)
	if err != nil {
		return fmt.Errorf("server not started: %w", err)
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
