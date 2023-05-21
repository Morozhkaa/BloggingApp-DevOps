package models

import "fmt"

var (
	ErrBadRequest        = fmt.Errorf("not filled with required parameters / invalid operation ID / changing invalid fields") // 400
	ErrForbidden         = fmt.Errorf("forbidden: wrong password")
	ErrTokenExpired      = fmt.Errorf("token expired")
	ErrNotFound          = fmt.Errorf("user not found")
	ErrGenerateToken     = fmt.Errorf("generate token failed")
	ErrBadAuth           = fmt.Errorf("required fields for authorization are not filled")
	ErrUserAlreadyExists = fmt.Errorf("user with this email already exists")
)
