package repository

import (
	"github.com/my-way-teams/my_way_backend/entity"
	"gorm.io/gorm"
)

type PlaceRepository interface {
	InsertPlace(c entity.Place) entity.Place
	UpdatePlace(c entity.Place) entity.Place
	DeletePlace(c entity.Place)
	AllPlace(cityID uint64) []entity.Place
	AllPlaces() []entity.Place
	FindPlaceByID(placeID uint64) entity.Place
}

type placeConnection struct {
	connection *gorm.DB
}

func NewPlaceRepository(dbConn *gorm.DB) PlaceRepository {
	return &placeConnection{
		connection: dbConn,
	}
}

func (db *placeConnection) InsertPlace(c entity.Place) entity.Place {
	db.connection.Save(&c)
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Find(&c)
	return c
}

func (db *placeConnection) UpdatePlace(c entity.Place) entity.Place {
	db.connection.Save(&c)
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Find(&c)
	return c
}

func (db *placeConnection) DeletePlace(c entity.Place) {
	db.connection.Delete(&c)
}

func (db *placeConnection) FindPlaceByID(placeID uint64) entity.Place {
	var place entity.Place
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Find(&place, placeID)
	return place
}

func (db *placeConnection) AllPlace(cityID uint64) []entity.Place {
	var places []entity.Place
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Where("city_id = ?", cityID).Find(&places)
	return places
}

func (db *placeConnection) AllPlaces() []entity.Place {
	var places []entity.Place
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Find(&places)
	return places
}