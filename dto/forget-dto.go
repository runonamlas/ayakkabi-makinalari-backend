package dto

type ForgetDTO struct {
	Email string `json:"email" form:"email" binding:"required"`
}
