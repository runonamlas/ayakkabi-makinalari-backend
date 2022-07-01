package repository

import (
	"github.com/my-way-teams/my_way_backend/entity"
	"gorm.io/gorm"
)

type CountryRepository interface {
	InsertCountry(c entity.Country) entity.Country
	UpdateCountry(c entity.Country) entity.Country
	DeleteCountry(c entity.Country)
	AllCountry() []entity.Country
	FindCountryByID(countryID uint64) entity.Country
}

type countryConnection struct {
	connection *gorm.DB
}

func NewCountryRepository(dbConn *gorm.DB) CountryRepository {
	return &countryConnection{
		connection: dbConn,
	}
}

func (db *countryConnection) InsertCountry(c entity.Country) entity.Country {
	db.connection.Save(&c)
	return c
}

func (db *countryConnection) UpdateCountry(c entity.Country) entity.Country {
	db.connection.Save(&c)
	return c
}

func (db *countryConnection) DeleteCountry(c entity.Country) {
	db.connection.Delete(&c)
}

func (db *countryConnection) FindCountryByID(countryID uint64) entity.Country {
	var country entity.Country
	db.connection.Find(&country, countryID)
	return country
}

func (db *countryConnection) AllCountry() []entity.Country {
	var countries []entity.Country
	db.connection.Find(&countries)
	return countries
}