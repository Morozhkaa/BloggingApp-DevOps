package http

import (
	"net/http"
	"user-app/internal/domain/models"

	"github.com/gin-gonic/gin"
	"github.com/juju/zaputil/zapctx"
	"go.uber.org/zap"
)

func (a *Adapter) login(ctx *gin.Context) {

	log := zap.L()
	log.Info("start login")
	var user models.UserAuth
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Info("error in login (BindJSON)")
		a.ErrorHandler(ctx, models.ErrBadRequest)
		return
	}

	ctx.Request = ctx.Request.WithContext(zapctx.WithFields(ctx.Request.Context(),
		zap.String("login", user.Login),
	))

	accessToken, refreshToken, err := a.auth.Login(ctx, user.Login, user.Password)
	if err != nil {
		log.Info("error in login usecase")
		a.ErrorHandler(ctx, err)
		return
	}

	cookie_access := &http.Cookie{
		Name:   "access_token",
		Value:  accessToken,
		Secure: false,
	}
	cookie_refresh := &http.Cookie{
		Name:   "refresh_token",
		Value:  refreshToken,
		Secure: false,
	}
	http.SetCookie(ctx.Writer, cookie_access)
	http.SetCookie(ctx.Writer, cookie_refresh)
	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
	log.Info("end login")
}

func (a *Adapter) verify(ctx *gin.Context) {

	access, err := ctx.Cookie("access_token")
	if err != nil {
		return
	}
	refresh, err := ctx.Cookie("refresh_token")
	if err != nil {
		return
	}

	resp, err := a.auth.Verify(ctx, access, refresh)
	if err != nil {
		a.ErrorHandler(ctx, err)
		return
	}

	cookie_access := &http.Cookie{
		Name:   "access_token",
		Value:  resp.AccessToken,
		Secure: false,
	}
	cookie_refresh := &http.Cookie{
		Name:   "refresh_token",
		Value:  resp.RefreshToken,
		Secure: false,
	}
	http.SetCookie(ctx.Writer, cookie_access)
	http.SetCookie(ctx.Writer, cookie_refresh)

	ctx.JSON(http.StatusOK, gin.H{
		"login": resp.Login,
		"email": resp.Email,
		"role":  resp.Role,
	})
}

func (a *Adapter) signup(ctx *gin.Context) {

	log := zap.L()
	log.Info("start signup")

	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Info("error in signup (BindJSON)")
		a.ErrorHandler(ctx, models.ErrBadRequest)
		return
	}

	err = a.auth.Signup(ctx, user)
	if err != nil {
		log.Info("error after signup usecase")
		a.ErrorHandler(ctx, err)
		return
	}
	log.Info("end signup")
}
