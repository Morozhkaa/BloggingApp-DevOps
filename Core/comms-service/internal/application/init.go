package application

import (
	"comm-service/internal/adapters/db"
	"comm-service/internal/adapters/http"
	"comm-service/internal/domain/usecases"
	"context"
	"fmt"
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

	userStorage, err := db.New(context.Background(), app.opts.DB_url)
	if err != nil {
		return fmt.Errorf("create comm storage failed: %w", err)
	}

	comm_service := usecases.New(userStorage)

	optsAdapter := http.AdapterOptions{HTTP_port: app.opts.HTTP_port, IsProd: app.opts.IsProd, AuthURL: app.opts.AuthURL}
	s, err := http.New(comm_service, optsAdapter)
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
