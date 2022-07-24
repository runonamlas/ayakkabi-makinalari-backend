package entity

import (
	"time"
)

type Product struct {
	ID           uint64          `gorm:"primary_key;auto_increment" json:"id"`
	Name         string          `gorm:"type:varchar(255)" json:"name"`
	Images       string          `gorm:"type:varchar(1000)" json:"images"`
	Brand        string          `gorm:"type:varchar(255)" json:"brand"`
	Used         string          `gorm:"type:varchar(20)" json:"used"`
	Price        string          `gorm:"type:varchar(255)" json:"price"`
	PriceUnit    uint8           `json:"priceUnit"`
	Vitrin       uint8           `json:"vitrin"`
	ClickProduct uint64          `gorm:"type:uint;default:0" json:"clickProduct"`
	CategoryID   uint64          `gorm:"" json:"-"`
	UserID       uint64          `gorm:"" json:"-"`
	User         User            `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"users"`
	Category     ProductCategory `gorm:"foreignKey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"categories"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
}
