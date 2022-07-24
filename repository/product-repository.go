package repository

import (
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProduct(c entity.Product) entity.Product
	UpdateProduct(c entity.Product) entity.Product
	DeleteProduct(c entity.Product)
	AllProduct(cityID uint64) []entity.Product
	AllProducts() []entity.Product
	FindProductByID(productID uint64) entity.Product
	FindProductByCategory(categoryID uint64) []entity.Product
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(dbConn *gorm.DB) ProductRepository {
	return &productConnection{
		connection: dbConn,
	}
}

func (db *productConnection) InsertProduct(c entity.Product) entity.Product {
	db.connection.Save(&c)
	db.connection.Preload("User").Preload("Category").Find(&c)
	return c
}

func (db *productConnection) UpdateProduct(c entity.Product) entity.Product {
	db.connection.Save(&c)
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Find(&c)
	return c
}

func (db *productConnection) DeleteProduct(c entity.Product) {
	db.connection.Delete(&c)
}

func (db *productConnection) FindProductByID(productID uint64) entity.Product {
	var product entity.Product
	db.connection.Preload("User").Preload("Category").Find(&product, productID)
	product.ClickProduct = product.ClickProduct + 1
	db.connection.Save(&product)
	return product
}

func (db *productConnection) FindProductByCategory(categoryID uint64) []entity.Product {
	var products []entity.Product
	db.connection.Preload("User").Preload("Category").Find(&products)
	return products
}

func (db *productConnection) AllProduct(cityID uint64) []entity.Product {
	var products []entity.Product
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Where("city_id = ?", cityID).Find(&products)
	return products
}

func (db *productConnection) AllProducts() []entity.Product {
	var products []entity.Product
	db.connection.Order("ID DESC").Preload("User").Preload("Category").Find(&products)
	return products
}
