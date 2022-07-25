package dto

type ChangePasswordDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Password string `json:"password" form:"password" binding:"required"`
}
