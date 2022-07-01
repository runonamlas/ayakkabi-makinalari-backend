package dto

import "github.com/my-way-teams/my_way_backend/entity"

type RouteUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
	Image  string `json:"image" form:"image" binding:"required"`
	Places []entity.Place `json:"places,omitempty" form:"places,omitempty"`
	CityID uint64 `json:"cityID,omitempty" form:"cityID,omitempty"`
}

type RouteCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
	Image  string `json:"image" form:"image" binding:"required"`
	Places []entity.Place `json:"places,omitempty" form:"places,omitempty"`
	CityID uint64 `json:"cityID,omitempty" form:"cityID,omitempty"`
}