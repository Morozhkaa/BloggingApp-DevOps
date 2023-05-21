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

type Post struct {
	ID          uuid.UUID `json:"post_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Created_at  Time      `json:"created_at"`
	Updated_at  Time      `json:"updated_at"`
	Author      Author    `json:"author"`
}

type Author struct {
	Email strfmt.Email `json:"email"`
	Login string       `json:"login"`
}

type NewPostDescription struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
