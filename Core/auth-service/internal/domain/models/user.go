package models

type UserRole string

const (
	RoleAdmin     UserRole = "ADMIN"
	RoleModerator UserRole = "MODERATOR"
	RoleMember    UserRole = "MEMBER"
	RoleGuest     UserRole = "GUEST"
)

type User struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Role     UserRole `json:"role"`
}

type UserAuth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type VerifyResponse struct {
	AccessToken  string
	RefreshToken string
	Login        string
	Email        string
	Role         UserRole
}
