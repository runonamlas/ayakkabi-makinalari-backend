package dto

type CountryUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
	Flag string `json:"flag" form:"flag" binding:"required"`
}

type CountryCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
	Flag string `json:"flag" form:"flag" binding:"required"`
}