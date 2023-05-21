package http

import (
	"comm-service/internal/domain/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/zaputil/zapctx"
)

func (a *Adapter) ErrorHandler(ctx *gin.Context, err error) {

	l := zapctx.Logger(ctx)
	l.Sugar().Errorf("request failed: %s", err.Error())

	switch {
	case errors.Is(err, models.ErrForbidden):
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
	case errors.Is(err, models.ErrBadAuth),
		errors.Is(err, models.ErrBadRequest):
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	case errors.Is(err, models.ErrNotFound):
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
