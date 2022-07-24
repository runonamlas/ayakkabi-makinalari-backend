package entity

import (
	"time"
)

type Message struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	MessageText string    `gorm:"type:varchar(255)" json:"name"`
	OwnerID     uint64    `gorm:"" json:"-"`
	Owner       User      `gorm:"foreignKey:OwnerID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"owner"`
	ProductID   uint64    `gorm:"" json:"-"`
	Product     Product   `gorm:"foreignKey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"product"`
	UserID      uint64    `gorm:"" json:"-"`
	User        User      `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	CreatedAt   time.Time `json:"createdAt"`
}
