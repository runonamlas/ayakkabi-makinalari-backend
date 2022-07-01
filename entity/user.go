package entity

type User struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Username string `uniqueIndex;gorm:"type:varchar(255)" json:"username"`
	CallNumber string `uniqueIndex;gorm:"type:varchar(25)" json:"callNumber"`
	Email string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token string `gorm:"-" json:"token,omitempty"`
	Places []Place `gorm:"many2many:user_place;" json:"places"`
	Routes []Route `gorm:"many2many:user_route;" json:"routes"`
}