package schemas

import "time"

type ForgotPassword struct {
	Email string `json:"email" binding:"required"`
}

type ResetPassword struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
	Email       string `json:"email" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required" example:"nandarusfikri@gmail.com"`
	Password string `json:"password" binding:"required" example:"Password1!"`
}
type LoginResponse struct {
	Id          uint64    `json:"id,omitempty"`
	Name        string    `json:"name" `
	Email       string    `json:"email" `
	AvatarPath  string    `json:"avatar_path" `
	AccessToken string    `json:"access_token"`
	ExpiredAt   time.Time `json:"expired_at"`
}
