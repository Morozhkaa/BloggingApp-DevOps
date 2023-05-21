package models

import "fmt"

var (
	ErrBadRequest = fmt.Errorf("not filled with required parameters / invalid operation ID / changing invalid fields") // 400
	ErrForbidden  = fmt.Errorf("forbidden")
	ErrNotFound   = fmt.Errorf("post not found")
	ErrBadAuth    = fmt.Errorf("required fields for authorization are not filled")
)
