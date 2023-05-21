package http

import (
	"errors"
	"fmt"
	"net/http"
	"post-service/internal/domain/models"
	"post-service/internal/ports"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func authMiddleware(a ports.AuthAdapter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("role", models.RoleGuest)
		err := a.Verify(ctx)
		if err != nil {
			switch {
			case errors.Is(err, models.ErrForbidden):
				ctx.JSON(http.StatusForbidden, gin.H{
					"error": err.Error(),
				})
			default:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
		}
		ctx.Next()
	}
}

func (a *Adapter) GetIdFromPath(ctx *gin.Context) (post_id uuid.UUID) {
	if ctx.Param("postId") == ":postId" { // if path-parameter is not set (this is how it works in Postman)
		a.ErrorHandler(ctx, models.ErrBadRequest)
		return uuid.Nil
	}
	post_id, err := uuid.Parse(ctx.Param("postId")) // may be invalid UUID length or formatis
	if err != nil {
		a.ErrorHandler(ctx, models.ErrBadRequest)
		return uuid.Nil
	}
	return post_id
}

func (a *Adapter) CreatePost(ctx *gin.Context) {
	// Публиковать посты могут только авторизованные пользователи
	role := ctx.MustGet("role").(models.UserRole)
	if role == models.RoleGuest {
		a.ErrorHandler(ctx, models.ErrForbidden)
		return
	}
	email := strfmt.Email(ctx.MustGet("email").(string))
	login := ctx.MustGet("login").(string)

	var data models.NewPostDescription
	err := ctx.BindJSON(&data)

	if err != nil {
		a.ErrorHandler(ctx, models.ErrBadRequest)
		return
	}

	t, post_id, err := a.PostService.Create(ctx, data, login, email)
	if err != nil {
		a.ErrorHandler(ctx, err)
		return
	}
	var created_at, updated_at models.Time
	created_at.Date = fmt.Sprintf("%04d/%02d/%02d", t.Year(), int(t.Month()), t.Day())
	created_at.Time = fmt.Sprintf("%02d:%02d:%02d", (t.Hour()+3)%24, t.Minute(), t.Second())
	updated_at = created_at

	ctx.JSON(http.StatusCreated, gin.H{
		"post_id":    post_id,
		"login":      login,
		"email":      email,
		"created_at": created_at,
		"updated_at": updated_at,
	})
}

func (a *Adapter) GetPost(ctx *gin.Context) {
	post_id := a.GetIdFromPath(ctx)
	if post_id == uuid.Nil {
		return
	}
	var data models.Post
	ctx.BindJSON(&data)

	post, err := a.PostService.Get(ctx, post_id)
	if err != nil {
		a.ErrorHandler(ctx, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, post)
}

func (a *Adapter) GetAllPosts(ctx *gin.Context) {
	posts, err := a.PostService.GetAllPosts(ctx)
	if err != nil {
		a.ErrorHandler(ctx, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, posts)
}

func (a *Adapter) UpdatePost(ctx *gin.Context) {
	post_id := a.GetIdFromPath(ctx)
	if post_id == uuid.Nil {
		return
	}
	email := strfmt.Email(ctx.MustGet("email").(string))
	var data models.NewPostDescription
	ctx.BindJSON(&data)
	role := ctx.MustGet("role").(models.UserRole)

	post, err := a.PostService.Update(ctx, post_id, email, data, role)
	if err != nil {
		a.ErrorHandler(ctx, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, post)
}

func (a *Adapter) DeletePost(ctx *gin.Context) {

	post_id := a.GetIdFromPath(ctx)
	if post_id == uuid.Nil {
		return
	}
	email := strfmt.Email(ctx.MustGet("email").(string))
	role := ctx.MustGet("role").(models.UserRole)
	log := zap.L()
	log.Info("IN delete_post (role): " + string(role))

	err := a.PostService.Delete(ctx, post_id, email, role)
	if err != nil {
		a.ErrorHandler(ctx, err)
		return
	}
	ctx.Status(http.StatusOK)
}
