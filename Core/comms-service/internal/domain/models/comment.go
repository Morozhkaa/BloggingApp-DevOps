package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type UserRole string

const (
	RoleAdmin     UserRole = "ADMIN"
	RoleModerator UserRole = "MODERATOR"
	RoleMember    UserRole = "MEMBER"
	RoleGuest     UserRole = "GUEST"
)

type Time struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

type Comment struct {
	ID         uuid.UUID `json:"comment_id"`
	Post_id    uuid.UUID `json:"post_id"`
	Text       string    `json:"text"`
	Created_at Time      `json:"created_at"`
	Updated_at Time      `json:"updated_at"`
	Commenter  Commenter `json:"commenter"`
}

type Commenter struct {
	Email strfmt.Email `json:"email"`
	Login string       `json:"login"`
}

type NewCommDescription struct {
	Text string `json:"text"`
}
