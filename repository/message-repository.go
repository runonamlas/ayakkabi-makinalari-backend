package repository

import (
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"gorm.io/gorm"
)

type MessageRepository interface {
	InsertMessage(c entity.Message) entity.Message
	UpdateMessage(c entity.Message) entity.Message
	DeleteMessage(c entity.Message)
	AllMessage(cityID uint64) []entity.Message
	AllMessages() []entity.Message
	FindMessageByID(messageID uint64) entity.Message
}

type messageConnection struct {
	connection *gorm.DB
}

func NewMessageRepository(dbConn *gorm.DB) MessageRepository {
	return &messageConnection{
		connection: dbConn,
	}
}

func (db *messageConnection) InsertMessage(c entity.Message) entity.Message {
	db.connection.Save(&c)
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Find(&c)
	return c
}

func (db *messageConnection) UpdateMessage(c entity.Message) entity.Message {
	db.connection.Save(&c)
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Find(&c)
	return c
}

func (db *messageConnection) DeleteMessage(c entity.Message) {
	db.connection.Delete(&c)
}

func (db *messageConnection) FindMessageByID(messageID uint64) entity.Message {
	var message entity.Message
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Find(&message, messageID)
	return message
}

func (db *messageConnection) AllMessage(cityID uint64) []entity.Message {
	var messages []entity.Message
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Where("city_id = ?", cityID).Find(&messages)
	return messages
}

func (db *messageConnection) AllMessages() []entity.Message {
	var messages []entity.Message
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Find(&messages)
	return messages
}
