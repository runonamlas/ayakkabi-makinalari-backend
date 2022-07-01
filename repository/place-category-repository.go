package repository

import (
	"github.com/my-way-teams/my_way_backend/entity"
	"gorm.io/gorm"
)

type PlaceCategoryRepository interface {
	InsertCategory(c entity.PlaceCategory) entity.PlaceCategory
	UpdateCategory(c entity.PlaceCategory) entity.PlaceCategory
	DeleteCategory(c entity.PlaceCategory)
	AllCategory() []entity.PlaceCategory
	FindCategoryByID(categoryID uint64) entity.PlaceCategory
}

type placeCategoryConnection struct {
	connection *gorm.DB
}

func NewPlaceCategoryRepository(dbConn *gorm.DB) PlaceCategoryRepository {
	return &placeCategoryConnection{
		connection: dbConn,
	}
}

func (db *placeCategoryConnection) InsertCategory(c entity.PlaceCategory) entity.PlaceCategory {
	db.connection.Save(&c)
	return c
}

func (db *placeCategoryConnection) UpdateCategory(c entity.PlaceCategory) entity.PlaceCategory {
	db.connection.Save(&c)
	return c
}

func (db *placeCategoryConnection) DeleteCategory(c entity.PlaceCategory) {
	db.connection.Delete(&c)
}

func (db *placeCategoryConnection) FindCategoryByID(categoryID uint64) entity.PlaceCategory {
	var category entity.PlaceCategory
	db.connection.Find(&category, categoryID)
	return category
}

func (db *placeCategoryConnection) AllCategory() []entity.PlaceCategory {
	var categories []entity.PlaceCategory
	db.connection.Find(&categories)
	return categories
}
