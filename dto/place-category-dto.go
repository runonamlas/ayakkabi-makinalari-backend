package dto

type PlaceCategoryUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
	Icon string `json:"icon" form:"icon" binding:"required"`
}

type PlaceCategoryCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
	Icon string `json:"icon" form:"icon" binding:"required"`
}