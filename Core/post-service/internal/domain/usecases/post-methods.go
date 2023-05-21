package usecases

import (
	"context"
	"post-service/internal/adapters/db"
	"post-service/internal/domain/models"
	"post-service/internal/ports"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type Post struct {
	postStorage ports.PostStorage
}

func New(postStorage ports.PostStorage) *Post {
	return &Post{
		postStorage: postStorage,
	}
}

func (a *Post) Create(ctx context.Context, postReq models.NewPostDescription, login string, email strfmt.Email) (created_at time.Time, post_id uuid.UUID, err error) {
	author := models.Author{
		Email: email,
		Login: login,
	}
	post := models.Post{
		ID:          uuid.New(),
		Title:       postReq.Title,
		Description: postReq.Description,
		Content:     postReq.Content,
		Author:      author,
	}

	created_at, err = a.postStorage.Save(ctx, post)

	if err != nil {
		return created_at, post.ID, err
	}
	return created_at, post.ID, nil
}

func (a *Post) Get(ctx context.Context, id uuid.UUID) (post models.Post, err error) {
	post, err = a.postStorage.Get(ctx, id)
	return post, err
}

func (a *Post) GetAllPosts(ctx context.Context) (posts []models.Post, err error) {
	posts, err = a.postStorage.GetAll(ctx)
	return posts, err
}

func (a *Post) Update(ctx context.Context, id uuid.UUID, email strfmt.Email,
	data models.NewPostDescription, role models.UserRole) (post models.Post, err error) {
	post, err = a.postStorage.Get(ctx, id)
	if err != nil {
		return post, err
	}
	if post.Author.Email != email && role != models.RoleAdmin && role != models.RoleModerator {
		return post, models.ErrForbidden
	}
	post.Title = data.Title
	post.Description = data.Description
	post.Content = data.Content
	time, err := a.postStorage.Update(ctx, id, data)
	db.ParseTime(&post.Updated_at, time)
	return post, err
}

func (a *Post) Delete(ctx context.Context, id uuid.UUID, email strfmt.Email, role models.UserRole) (err error) {
	err = a.postStorage.Delete(ctx, id, email, role)
	return err
}
