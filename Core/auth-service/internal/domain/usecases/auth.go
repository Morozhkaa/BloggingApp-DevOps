package usecases

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"time"
	"user-app/internal/config"
	"user-app/internal/domain/models"
	"user-app/internal/ports"

	jwt4 "github.com/golang-jwt/jwt/v4"
	"github.com/xdg-go/pbkdf2"
	"go.uber.org/zap"
)

type Auth struct {
	userStorage ports.UserStorage
}

func New(userStorage ports.UserStorage) *Auth {
	return &Auth{
		userStorage: userStorage,
	}
}

const accessToken_expiration_time = 1 * time.Minute
const refreshToken_expiration_time = 60 * time.Minute

func (a *Auth) EncodePassword(password string) string {
	cfg, _ := config.GetConfig()
	dk := pbkdf2.Key([]byte(password), []byte(cfg.Salt), 1000, 128, sha1.New)
	return base64.StdEncoding.EncodeToString([]byte(dk))
}

func (a *Auth) generateToken(login string, email string, role models.UserRole, expiredIn time.Duration) (token string, err error) {
	cfg, _ := config.GetConfig()

	now := time.Now()
	claims := &jwt4.RegisteredClaims{
		ExpiresAt: jwt4.NewNumericDate(now.Add(expiredIn)),
		Issuer:    login,
		Subject:   email,
		ID:        string(role),
	}

	t := jwt4.NewWithClaims(jwt4.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", models.ErrGenerateToken
	}
	return token, nil
}

func (a *Auth) verifyToken(tokenString string) (login, email string, role models.UserRole, err error) {
	cfg, _ := config.GetConfig()
	var claims jwt4.RegisteredClaims
	token, err := jwt4.ParseWithClaims(tokenString, &claims, func(token *jwt4.Token) (interface{}, error) {
		return []byte(cfg.Secret), nil
	})

	if !token.Valid {
		return "", "", "", fmt.Errorf("parse token unexpected error: %w", err)
	}

	return claims.Issuer, claims.Subject, models.UserRole(claims.ID), nil
}

func (a *Auth) Login(ctx context.Context, login, password string) (access, refresh string, err error) {

	log := zap.L()
	log.Info("start login usecase")
	user, err := a.userStorage.GetUser(ctx, login)
	if err != nil {
		log.Info("error in login usecase (GetUser)")
		return "", "", err
	}
	if user.Password != a.EncodePassword(password) {
		log.Info("error in login usecase (WrongPassword)")
		return "", "", models.ErrForbidden
	}

	access, err = a.generateToken(login, user.Email, user.Role, accessToken_expiration_time)
	if err != nil {
		log.Info("error in login usecase (generate access Token)")
		return "", "", err
	}
	refresh, err = a.generateToken(login, user.Email, user.Role, refreshToken_expiration_time)
	if err != nil {
		log.Info("error in login usecase (generate refresh Token)")
		return "", "", err
	}
	log.Info("end login usecase")

	return access, refresh, nil
}

func (a *Auth) Verify(ctx context.Context, access string, refresh string) (r models.VerifyResponse, err error) {

	r.Login, r.Email, r.Role, err = a.verifyToken(access)
	if err == nil {
		r.AccessToken = access
		r.RefreshToken = refresh
		return r, nil
	}

	r.Login, r.Email, r.Role, err = a.verifyToken(refresh)
	if err == nil {
		r.AccessToken, err = a.generateToken(r.Login, r.Email, r.Role, accessToken_expiration_time)
		if err != nil {
			return r, err
		}
		r.RefreshToken, err = a.generateToken(r.Login, r.Email, r.Role, refreshToken_expiration_time)
		if err != nil {
			return r, err
		}
		return r, nil
	}

	return r, models.ErrTokenExpired
}

func (a *Auth) Signup(ctx context.Context, user models.User) (err error) {
	user.Password = a.EncodePassword(user.Password)
	err = a.userStorage.SaveUser(ctx, user)
	return err
}
