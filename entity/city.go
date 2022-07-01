package entity

type City struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"uniqueIndex;type:varchar(255)" json:"name"`
	Image  string `gorm:"type:varchar(255)" json:"image"`
	CountryID uint64 `gorm:"not null" json:"-"`
	Country Country `gorm:"foreignKey:CountryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"country"`
}