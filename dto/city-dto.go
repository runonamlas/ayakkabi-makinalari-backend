package dto

type CityUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
	Image  string `json:"image" form:"image" binding:"required"`
	CountryID uint64 `json:"countryID,omitempty" form:"countryID,omitempty"`
}

type CityCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
	Image  string `json:"image" form:"image" binding:"required"`
	CountryID uint64 `json:"countryID,omitempty" form:"countryID,omitempty"`
}
