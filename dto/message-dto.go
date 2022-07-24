package dto

type MessageUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	MessageText string `json:"message" form:"message" binding:"required"`
	Owner       string `json:"owner" form:"owner" binding:"required"`
	ProductID   uint64 `json:"productID,omitempty" form:"productID,omitempty"`
	UserID      uint64 `json:"userID,omitempty" form:"userID,omitempty"`
}

type MessageCreateDTO struct {
	MessageText string `json:"message" form:"message" binding:"required"`
	UserID      uint64 `json:"userID,omitempty" form:"userID,omitemptiy" binding:"required"`
	ProductID   uint64 `json:"productID,omitempty" form:"productID,omitempty"`
	OwnerID     uint64
}
