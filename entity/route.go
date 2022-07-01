package entity

type Route struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"uniqueIndex;type:varchar(255)" json:"name"`
	Image  string `gorm:"type:varchar(255)" json:"image"`
	CityID uint64 `gorm:"not null" json:"-"`
	City City `gorm:"foreignKey:CityID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"city"`
	Places []Place `gorm:"many2many:route_place;" json:"places"`
	Users []User `gorm:"many2many:user_route;" json:"users"`
}
