package repository

import (
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"gorm.io/gorm"
)

type ProductCategoryRepository interface {
	InsertCategory(c entity.ProductCategory) entity.ProductCategory
	UpdateCategory(c entity.ProductCategory) entity.ProductCategory
	DeleteCategory(c entity.ProductCategory)
	AllCategory() []entity.ProductCategory
	FindCategoryByID(categoryID uint64) entity.ProductCategory
}

type productCategoryConnection struct {
	connection *gorm.DB
}

func NewProductCategoryRepository(dbConn *gorm.DB) ProductCategoryRepository {
	return &productCategoryConnection{
		connection: dbConn,
	}
}

func (db *productCategoryConnection) InsertCategory(c entity.ProductCategory) entity.ProductCategory {
	db.connection.Save(&c)
	return c
}

func (db *productCategoryConnection) UpdateCategory(c entity.ProductCategory) entity.ProductCategory {
	db.connection.Save(&c)
	return c
}

func (db *productCategoryConnection) DeleteCategory(c entity.ProductCategory) {
	db.connection.Delete(&c)
}

func (db *productCategoryConnection) FindCategoryByID(categoryID uint64) entity.ProductCategory {
	var category entity.ProductCategory
	db.connection.Preload("Products.User").Find(&category, categoryID)
	return category
}

func (db *productCategoryConnection) AllCategory() []entity.ProductCategory {
	var categories []entity.ProductCategory
	db.connection.Preload("Products").Find(&categories)
	return categories
}
