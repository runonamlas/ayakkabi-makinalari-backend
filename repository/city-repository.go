package repository

import (
	"github.com/my-way-teams/my_way_backend/entity"
	"gorm.io/gorm"
)

type CityRepository interface {
	InsertCity(c entity.City) entity.City
	UpdateCity(c entity.City) entity.City
	DeleteCity(c entity.City)
	AllCity(countryID uint64) []entity.City
	AllCities() []entity.City
	FindCityByID(cityID uint64) entity.City
}

type cityConnection struct {
	connection *gorm.DB
}

func NewCityRepository(dbConn *gorm.DB) CityRepository {
	return &cityConnection{
		connection: dbConn,
	}
}

func (db *cityConnection) InsertCity(c entity.City) entity.City {
	db.connection.Save(&c)
	db.connection.Preload("Country").Find(&c)
	return c
}

func (db *cityConnection) UpdateCity(c entity.City) entity.City {
	db.connection.Save(&c)
	db.connection.Preload("Country").Find(&c)
	return c
}

func (db *cityConnection) DeleteCity(c entity.City) {
	db.connection.Delete(&c)
}

func (db *cityConnection) FindCityByID(cityID uint64) entity.City {
	var city entity.City
	db.connection.Preload("Country").Find(&city, cityID)
	return city
}

func (db *cityConnection) AllCity(countryID uint64) []entity.City {
	var cities []entity.City
	db.connection.Preload("Country").Where("country_id = ?", countryID).Find(&cities)
	return cities
}

func (db *cityConnection) AllCities() []entity.City {
	var cities []entity.City
	db.connection.Preload("Country").Find(&cities)
	return cities
}