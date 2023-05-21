package ports

import (
	"context"
	"post-service/internal/domain/models"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type PostStorage interface {
	GetAll(ctx context.Context) ([]models.Post, error)
	Get(ctx context.Context, id uuid.UUID) (models.Post, error)
	Delete(ctx context.Context, id uuid.UUID, email strfmt.Email, role models.UserRole) error
	Save(ctx context.Context, post models.Post) (time.Time, error)
	Update(ctx context.Context, id uuid.UUID, data models.NewPostDescription) (time.Time, error)
}
