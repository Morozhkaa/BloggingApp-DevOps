package ports

import (
	"context"
	"post-service/internal/domain/models"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type Post interface {
	Create(ctx context.Context, post models.NewPostDescription, login string, user strfmt.Email) (time.Time, uuid.UUID, error)
	Get(ctx context.Context, id uuid.UUID) (models.Post, error)
	GetAllPosts(ctx context.Context) ([]models.Post, error)
	Update(ctx context.Context, id uuid.UUID, user strfmt.Email, data models.NewPostDescription, role models.UserRole) (models.Post, error)
	Delete(ctx context.Context, id uuid.UUID, user strfmt.Email, role models.UserRole) error
}
