package ports

import (
	"comm-service/internal/domain/models"
	"context"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type Comment interface {
	Create(ctx context.Context, comment models.NewCommDescription, post_id uuid.UUID, login string, user strfmt.Email) (time.Time, uuid.UUID, error)
	Get(ctx context.Context, comm_id uuid.UUID) (models.Comment, error)
	GetAllComments(ctx context.Context, post_id uuid.UUID) ([]models.Comment, error)
	Update(ctx context.Context, comm_id uuid.UUID, user strfmt.Email, data models.NewCommDescription, role models.UserRole) (models.Comment, error)
	Delete(ctx context.Context, comm_id uuid.UUID, user strfmt.Email, role models.UserRole) error
}
