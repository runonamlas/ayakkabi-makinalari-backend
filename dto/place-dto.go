package dto

type PlaceUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
	Image  string `json:"image" form:"image" binding:"required"`
	Desc  string `json:"desc" form:"desc" binding:"required"`
	Star float64 `json:"star" form:"star" binding:"required"`
	Lat float64 `json:"lat" form:"lat" binding:"required"`
	Long float64 `json:"long" form:"long" binding:"required"`
	CityID uint64 `json:"cityID,omitempty" form:"cityID,omitempty"`
	CategoryID uint64 `json:"categoryID,omitempty" form:"categoryID,omitempty"`
}

type PlaceCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
	Image  string `json:"image" form:"image" binding:"required"`
	Desc  string `json:"desc" form:"desc" binding:"required"`
	Star float64 `json:"star" form:"star" binding:"required"`
	Lat float64 `json:"lat" form:"lat" binding:"required"`
	Long float64 `json:"long" form:"long" binding:"required"`
	CityID uint64 `json:"cityID,omitempty" form:"cityID,omitempty"`
	CategoryID uint64 `json:"categoryID,omitempty" form:"categoryID,omitempty"`
}