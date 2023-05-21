package db

import (
	"context"
	"errors"
	"time"
	"user-app/internal/domain/models"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

type DBStorage struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, conn string) (*DBStorage, error) {
	time.Sleep(4 * time.Second)
	pool, err := pgxpool.Connect(ctx, conn)
	if err != nil {
		return nil, err
	}
	return &DBStorage{
		pool: pool,
	}, nil
}

func (db *DBStorage) GetUser(ctx context.Context, login string) (user models.User, err error) {
	const query = `
	SELECT email, login, password, role FROM users WHERE login = $1
	`
	err = db.pool.QueryRow(ctx, query, login).Scan(
		&user.Email,
		&user.Login,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, nil
		}
		return user, models.ErrNotFound
	}
	return user, nil
}

func (db *DBStorage) SaveUser(ctx context.Context, user models.User) (err error) {
	var count int
	const queryCheck = `
	SELECT COUNT(*) FROM users WHERE email = $1
	`
	err = db.pool.QueryRow(ctx, queryCheck, user.Email).Scan(
		&count,
	)
	if count != 0 {
		return models.ErrUserAlreadyExists
	}

	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	const query = `
	INSERT INTO users (login, password, email)
	VALUES ($1, $2, $3)
	`
	_, err = tx.Exec(ctx, query, user.Login, user.Password, user.Email)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	return err
}
