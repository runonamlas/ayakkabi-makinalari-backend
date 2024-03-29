package dto

type RegisterDTO struct {
	Username   string `json:"username" form:"username" binding:"required"`
	CallNumber string `json:"callNumber" form:"callNumber" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required,email"`
	Address    string `json:"address" form:"address" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
}
