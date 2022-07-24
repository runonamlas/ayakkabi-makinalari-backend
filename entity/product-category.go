package entity

type ProductCategory struct {
	ID       uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Name     string     `gorm:"uniqueIndex;type:varchar(255)" json:"name"`
	Products []*Product `gorm:"foreignKey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"products"`
}
