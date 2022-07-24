package dto

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Address  string `json:"address" form:"address"`
}
