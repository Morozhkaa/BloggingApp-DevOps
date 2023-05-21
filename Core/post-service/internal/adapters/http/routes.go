package http

import (
	"fmt"
	"post-service/logger"
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

	v1 := r.Group("/api/post-service/v1")
	{
		v1.GET("/getAll", a.GetAllPosts)
	}
	v1_auth := r.Group("/api/post-service/v1").Use(authMiddleware(a))
	{
		v1_auth.POST("/create", a.CreatePost)
		v1_auth.GET("/get/:postId", a.GetPost)
		v1_auth.POST("/update/:postId", a.UpdatePost)
		v1_auth.DELETE("/delete/:postId", a.DeletePost)
	}
	return nil
}
