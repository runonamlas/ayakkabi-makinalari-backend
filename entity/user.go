package entity

import "time"

type User struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Username     string    `gorm:"uniqueIndex;type:varchar(255)" json:"username"`
	CallNumber   string    `gorm:"uniqueIndex;type:varchar(25)" json:"callNumber"`
	Email        string    `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password     string    `gorm:"->;<-;not null" json:"-"`
	Token        string    `gorm:"-" json:"token,omitempty"`
	Address      string    `gorm:"type:varchar(255)" json:"address"`
	AccountType  uint8     `gorm:"default:0;" json:"accountType"`
	ClickProfile uint64    `gorm:"type:uint;default:0" json:"clickProfile"`
	Products     []Product `gorm:"many2many:user_product;" json:"products"`
	Messages     []Message `gorm:"many2many:user_message;" json:"messages"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
