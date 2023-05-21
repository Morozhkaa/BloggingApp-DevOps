package http

import (
	"comm-service/internal/domain/models"
	"comm-service/internal/ports"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
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

func (a *Adapter) GetIdFromPath(ctx *gin.Context, param string) (comm_id uuid.UUID) {
	if ctx.Param(param) == ":"+param { // if path-parameter is not set (this is how it works in Postman)
		a.ErrorHandler(ctx, models.ErrBadRequest)
		return uuid.Nil
	}
	id, err := uuid.Parse(ctx.Param(param)) // may be invalid UUID length or formatis
	if err != nil {
		a.ErrorHandler(ctx, models.ErrBadRequest)
		return uuid.Nil
	}
	return id
}

func (a *Adapter) CreateComment(ctx *gin.Context) {
	// Оставлять комментарии могут только авторизованные пользователи
	role := ctx.MustGet("role").(models.UserRole)
	if role == models.RoleGuest {
		a.ErrorHandler(ctx, models.ErrForbidden)
		return
	}

	post_id := a.GetIdFromPath(ctx, "postId")
	if post_id == uuid.Nil {
		return
	}
	email := strfmt.Email(ctx.MustGet("email").(string))
	login := ctx.MustGet("login").(string)

	var data models.NewCommDescription
	err := ctx.BindJSON(&data)
	if err != nil {
		a.ErrorHandler(ctx, models.ErrBadRequest)
		return
	}

	t, comm_id, err := a.CommService.Create(ctx, data, post_id, login, email)
	if err != nil {
		a.ErrorHandler(ctx, err)
		return
	}
	var created_at, updated_at models.Time
	created_at.Date = fmt.Sprintf("%04d/%02d/%02d", t.Year(), int(t.Month()), t.Day())
	created_at.Time = fmt.Sprintf("%02d:%02d:%02d", (t.Hour()+3)%24, t.Minute(), t.Second())
	updated_at = created_at

	ctx.JSON(http.StatusCreated, gin.H{
		"comm_id":    comm_id,
		"login":      login,
		"email":      email,
		"created_at": created_at,
		"updated_at": updated_at,
	})
}

func (a *Adapter) GetComment(ctx *gin.Context) {
	comm_id := a.GetIdFromPath(ctx, "commId")
	if comm_id == uuid.Nil {
		return
	}
	var data models.Comment
	ctx.BindJSON(&data)

	comment, err := a.CommService.Get(ctx, comm_id)
	if err != nil {
		a.ErrorHandler(ctx, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, comment)
}

func (a *Adapter) GetAllComments(ctx *gin.Context) {
	post_id := a.GetIdFromPath(ctx, "postId")
	if post_id == uuid.Nil {
		return
	}
	comments, err := a.CommService.GetAllComments(ctx, post_id)
	if err != nil {
		a.ErrorHandler(ctx, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, comments)
}

func (a *Adapter) UpdateComment(ctx *gin.Context) {
	comm_id := a.GetIdFromPath(ctx, "commId")
	if comm_id == uuid.Nil {
		return
	}
	email := strfmt.Email(ctx.MustGet("email").(string))
	var data models.NewCommDescription
	ctx.BindJSON(&data)
	role := ctx.MustGet("role").(models.UserRole)

	comment, err := a.CommService.Update(ctx, comm_id, email, data, role)
	if err != nil {
		a.ErrorHandler(ctx, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, comment)
}

func (a *Adapter) DeleteComment(ctx *gin.Context) {

	comm_id := a.GetIdFromPath(ctx, "commId")
	if comm_id == uuid.Nil {
		return
	}
	email := strfmt.Email(ctx.MustGet("email").(string))
	role := ctx.MustGet("role").(models.UserRole)

	err := a.CommService.Delete(ctx, comm_id, email, role)
	if err != nil {
		a.ErrorHandler(ctx, err)
		return
	}
	ctx.Status(http.StatusOK)
}
