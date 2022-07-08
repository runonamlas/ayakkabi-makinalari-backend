package dto

type ProductUpdateDTO struct {
	ID         uint64 `json:"id" form:"id" binding:"required"`
	Name       string `json:"name" form:"name" binding:"required"`
	Images     string `json:"images" form:"images" binding:"required"`
	Brand      string `json:"brand" form:"brand" binding:"required"`
	Used       bool   `json:"used" form:"used" binding:"required"`
	Price      string `json:"price" form:"price" binding:"required"`
	PriceUnit  uint8  `json:"priceUnit" form:"priceUnit" binding:"required"`
	CategoryID uint64 `json:"categoryID,omitempty" form:"categoryID,omitempty"`
}

type ProductCreateDTO struct {
	Name       string `json:"name" form:"name" binding:"required"`
	Images     string `json:"images" form:"images" binding:"required"`
	Brand      string `json:"brand" form:"brand" binding:"required"`
	Used       bool   `json:"used" form:"used" binding:"required"`
	Price      string `json:"price" form:"price" binding:"required"`
	PriceUnit  uint8  `json:"priceUnit" form:"priceUnit" binding:"required"`
	CategoryID uint64 `json:"categoryID,omitempty" form:"categoryID,omitempty"`
}
