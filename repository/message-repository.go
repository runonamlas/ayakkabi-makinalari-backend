package repository

import (
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"gorm.io/gorm"
)

type MessageRepository interface {
	InsertMessage(c entity.Message) entity.Message
	UpdateMessage(c entity.Message) entity.Message
	DeleteMessage(c entity.Message)
	AllMessage(userID uint64, ownerID uint64) []entity.Message
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
	db.connection.Preload("Owner").Preload("Product").Preload("User").Find(&c)
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

func (db *messageConnection) AllMessage(userID uint64, ownerID uint64) []entity.Message {
	var messages []entity.Message
	var messagess []entity.Message
	db.connection.Order("ID DESC").Preload("Owner").Preload("Product").Preload("User").Where("user_id = ? AND owner_id = ?", userID, ownerID).Find(&messages)
	db.connection.Order("ID DESC").Preload("Owner").Preload("Product").Preload("User").Where("user_id = ? AND owner_id = ?", ownerID, userID).Find(&messagess)
	messages = append(messages, messagess...)
	return messages
}

func (db *messageConnection) AllMessages() []entity.Message {
	var messages []entity.Message
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Find(&messages)
	return messages
}
