package dto

type AdminUpdateDTO struct {
	ID uint64 `json:"id" form:"id"`
	Username string `json:"username" form:"username" binding:"required"`
	Email string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}

type AdminRegisterDTO struct {
	Username string `json:"username" form:"username" binding:"required"`
	Email string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type AdminLoginDTO struct {
	Email string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}