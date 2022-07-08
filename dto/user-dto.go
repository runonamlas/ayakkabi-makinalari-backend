package dto

type UserUpdateDTO struct {
	ID         uint64 `json:"id" form:"id"`
	Username   string `json:"username" form:"username" binding:"required"`
	CallNumber string `json:"callNumber" form:"callNumber" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required,email"`
	Address    string `json:"address" form:"address" binding:"required"`
	Password   string `json:"password,omitempty" form:"password,omitempty"`
}
