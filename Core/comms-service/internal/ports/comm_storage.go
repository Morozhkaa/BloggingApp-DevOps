package ports

import (
	"comm-service/internal/domain/models"
	"context"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type CommStorage interface {
	GetAll(ctx context.Context, post_id uuid.UUID) ([]models.Comment, error)
	Get(ctx context.Context, comm_id uuid.UUID) (models.Comment, error)
	Delete(ctx context.Context, comm_id uuid.UUID, email strfmt.Email, role models.UserRole) error
	Save(ctx context.Context, comment models.Comment) (time.Time, error)
	Update(ctx context.Context, comm_id uuid.UUID, data models.NewCommDescription) (time.Time, error)
}
