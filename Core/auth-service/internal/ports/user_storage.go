package ports

import (
	"context"
	"user-app/internal/domain/models"
)

type UserStorage interface {
	GetUser(ctx context.Context, login string) (models.User, error)
	SaveUser(ctx context.Context, user models.User) error
}
