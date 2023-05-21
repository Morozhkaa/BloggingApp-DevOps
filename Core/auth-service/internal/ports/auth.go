package ports

import (
	"context"
	"user-app/internal/domain/models"
)

type Auth interface {
	Login(ctx context.Context, login, password string) (string, string, error)
	Verify(ctx context.Context, access, refresh string) (models.VerifyResponse, error)
	Signup(ctx context.Context, user models.User) error
}
