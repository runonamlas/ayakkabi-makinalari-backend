package dto

type ProductCategoryUpdateDTO struct {
	ID   uint64 `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

type ProductCategoryCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
}
