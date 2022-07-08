package entity

import "time"

type Product struct {
	ID         uint64          `gorm:"primary_key;auto_increment" json:"id"`
	Name       string          `gorm:"type:varchar(255)" json:"name"`
	Images     string          `gorm:"type:varchar(1000)" json:"images"`
	Brand      string          `gorm:"type:varchar(255)" json:"brand"`
	Used       bool            `gorm:"type:bool" json:"used"`
	Price      string          `gorm:"type:varchar(255)" json:"price"`
	PriceUnit  uint8           `json:"priceUnit"`
	CategoryID uint64          `gorm:"" json:"-"`
	Category   ProductCategory `gorm:"foreignKey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"category"`
	User       User            `gorm:"many2many:user_product;" json:"user"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
}
