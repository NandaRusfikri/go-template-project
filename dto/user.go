package dto

type UsersRequest struct {
	SearchText string `form:"search_text"`
	Page       int    `form:"page" example:"1"`
	Limit      int    `form:"limit" example:"10"`
	OrderField string `form:"order_field" example:"id|desc"`
	IsActive   *bool  `form:"is_active"`
}

type CheckEmail struct {
	Email string `json:"email" binding:"required" example:"nandarusfikri@gmail.com"`
}

type CheckUsername struct {
	Username string `json:"username" binding:"required" example:"nandarusfikri"`
}

type UsersResponse struct {
	Id         uint64 ` json:"id"`
	Name       string ` json:"name"`
	Email      string ` json:"email"`
	IsActive   bool   ` json:"is_active"`
	AvatarPath string ` json:"avatar_path"`
}
type UserInsert struct {
	Email    string `json:"email" binding:"required" example:"nandarusfikri@gmail.com"`
	Password string `json:"password" binding:"required" example:"Password1!"`
	Name     string `json:"name" binding:"required" example:"Nanda Rusfikri"`
}
type UserUpdate struct {
	Id       uint64 `json:"id" binding:"required" example:"1"`
	Email    string `json:"email" binding:"-" example:"nandarusfikri@gmail.com"`
	Name     string `json:"name" binding:"-" example:"Nanda Rusfikri"`
	Phone    string `json:"phone" binding:"-" example:"08123456789"`
	IsActive *bool  ` json:"is_active" example:"true"`
}

type UserDetail struct {
	Id         uint64 ` json:"id"`
	Name       string ` json:"name"`
	Email      string ` json:"email"`
	Username   string ` json:"username"`
	IsActive   bool   ` json:"is_active"`
	AvatarPath string ` json:"avatar_path"`
}

type ChangePassword struct {
	UserId      uint64 `json:"user_id" binding:"required" example:"1"`
	OldPassword string `json:"old_password" binding:"required" example:"Password1!"`
	NewPassword string `json:"new_password" binding:"required" example:"Password1!"`
}
