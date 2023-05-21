package usecases

import (
	"comm-service/internal/adapters/db"
	"comm-service/internal/domain/models"
	"comm-service/internal/ports"
	"context"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type Comment struct {
	commStorage ports.CommStorage
}

func New(commStorage ports.CommStorage) *Comment {
	return &Comment{
		commStorage: commStorage,
	}
}

func (a *Comment) Create(ctx context.Context, commReq models.NewCommDescription, post_id uuid.UUID, login string, email strfmt.Email) (created_at time.Time, comm_id uuid.UUID, err error) {
	commenter := models.Commenter{
		Email: email,
		Login: login,
	}
	comment := models.Comment{
		ID:        uuid.New(),
		Post_id:   post_id,
		Text:      commReq.Text,
		Commenter: commenter,
	}

	created_at, err = a.commStorage.Save(ctx, comment)

	if err != nil {
		return created_at, comment.ID, err
	}
	return created_at, comment.ID, nil
}

func (a *Comment) Get(ctx context.Context, comm_id uuid.UUID) (comment models.Comment, err error) {
	comment, err = a.commStorage.Get(ctx, comm_id)
	return comment, err
}

func (a *Comment) GetAllComments(ctx context.Context, post_id uuid.UUID) (comments []models.Comment, err error) {
	comments, err = a.commStorage.GetAll(ctx, post_id)
	return comments, err
}

func (a *Comment) Update(ctx context.Context, comm_id uuid.UUID, email strfmt.Email,
	data models.NewCommDescription, role models.UserRole) (comment models.Comment, err error) {
	comment, err = a.commStorage.Get(ctx, comm_id)
	if err != nil {
		return comment, err
	}
	if comment.Commenter.Email != email && role != models.RoleAdmin && role != models.RoleModerator {
		return comment, models.ErrForbidden
	}
	comment.Text = data.Text
	time, err := a.commStorage.Update(ctx, comm_id, data)
	db.ParseTime(&comment.Updated_at, time)
	return comment, err
}

func (a *Comment) Delete(ctx context.Context, id uuid.UUID, email strfmt.Email, role models.UserRole) (err error) {
	err = a.commStorage.Delete(ctx, id, email, role)
	return err
}
