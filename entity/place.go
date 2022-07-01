package entity

type Place struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"uniqueIndex;type:varchar(255)" json:"name"`
	Image  string `gorm:"type:varchar(255)" json:"image"`
	Desc string `gorm:"type:text" json:"desc"`
	Star float64 `gorm:"type:float" json:"star"`
	Lat float64 `gorm:"type:float" json:"lat"`
	Long float64 `gorm:"type:float" json:"long"`
	CityID uint64 `gorm:"not null" json:"-"`
	City City `gorm:"foreignKey:CityID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"city"`
	CategoryID uint64 `gorm:"" json:"-"`
	Category PlaceCategory `gorm:"foreignKey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"category"`
	Routes []Route `gorm:"many2many:route_place;" json:"routes"`
	Users []User `gorm:"many2many:user_place;" json:"users"`
}
