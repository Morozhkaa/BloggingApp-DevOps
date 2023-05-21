package http

import (
	"fmt"
	"time"
	"user-app/pkg/infra/logger"

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

	r.POST("/login", a.login)
	r.POST("/verify", a.verify)
	r.POST("/signup", a.signup)
	return nil
}
