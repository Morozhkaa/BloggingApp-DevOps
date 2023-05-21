package http

import (
	"comm-service/logger"
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/juju/zaputil/zapctx"
)

func initRouter(a *Adapter, r *gin.Engine, opts logger.LoggerOptions) error {
	l, err := logger.New(opts)
	if err != nil {
		return fmt.Errorf("logger initialization failed: %w", err)
	}

	r.Use(func(ctx *gin.Context) {
		lCtx := zapctx.WithLogger(ctx.Request.Context(), l)
		ctx.Request = ctx.Request.WithContext(lCtx)
	})
	r.Use(ginzap.Ginzap(l, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(l, true))

	v1 := r.Group("/api/comms-service/v1")
	{
		v1.GET("/getAll/:postId", a.GetAllComments)
	}
	v1_auth := r.Group("/api/comms-service/v1").Use(authMiddleware(a))
	{
		v1_auth.POST("/create/:postId", a.CreateComment)
		v1_auth.GET("/get/:commId", a.GetComment)
		v1_auth.POST("/update/:commId", a.UpdateComment)
		v1_auth.DELETE("/delete/:commId", a.DeleteComment)
	}
	return nil
}
