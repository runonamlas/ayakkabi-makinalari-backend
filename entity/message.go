package entity

import "time"

type Message struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	MessageText string    `gorm:"uniqueIndex;type:varchar(255)" json:"name"`
	Owner       string    `gorm:"type:varchar(255)" json:"owner"`
	ProductID   uint64    `gorm:"" json:"-"`
	Product     Product   `gorm:"foreignKey:ProductID;" json:"product"`
	User        User      `gorm:"many2many:user_message;" json:"user"`
	CreatedAt   time.Time `json:"createdAt"`
}
