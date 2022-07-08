package repository

import (
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"gorm.io/gorm"
)

type AdminRepository interface {
	InsertAdmin(admin entity.Admin) entity.Admin
	UpdateAdmin(admin entity.Admin) entity.Admin
	VerifyCredential(email string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	IsDuplicateUsername(username string) (tx *gorm.DB)
	FindByEmail(email string) entity.Admin
	FindByUsername(username string) entity.Admin
	ProfileAdmin(adminID string) entity.Admin
	Users() []entity.User
}

type adminConnection struct {
	connection *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminConnection{
		connection: db,
	}
}

func (db *adminConnection) InsertAdmin(admin entity.Admin) entity.Admin {
	admin.Password = hashAndSalt([]byte(admin.Password))
	db.connection.Save(&admin)
	return admin
}

func (db *adminConnection) UpdateAdmin(admin entity.Admin) entity.Admin {
	if admin.Password != "" {
		admin.Password = hashAndSalt([]byte(admin.Password))
	} else {
		var tempAdmin entity.Admin
		db.connection.Find(&tempAdmin, admin.ID)
		admin.Password = tempAdmin.Password
	}
	db.connection.Save(&admin)
	return admin
}

func (db *adminConnection) VerifyCredential(email string) interface{} {
	var admin entity.Admin
	res := db.connection.Where("email = ?", email).Take(&admin)
	resUsername := db.connection.Where("username = ?", email).Take(&admin)
	if res.Error == nil || resUsername.Error == nil {
		return admin
	}
	return nil
}

func (db *adminConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var admin entity.Admin
	return db.connection.Where("email = ?", email).Take(&admin)
}

func (db *adminConnection) IsDuplicateUsername(username string) (tx *gorm.DB) {
	var admin entity.Admin
	return db.connection.Where("username = ?", username).Take(&admin)
}

func (db *adminConnection) FindByEmail(email string) entity.Admin {
	var admin entity.Admin
	db.connection.Where("email = ?", email).Take(&admin)
	return admin
}

func (db *adminConnection) FindByUsername(username string) entity.Admin {
	var admin entity.Admin
	db.connection.Where("username = ?", username).Take(&admin)
	return admin
}

func (db *adminConnection) ProfileAdmin(adminID string) entity.Admin {
	var admin entity.Admin
	db.connection.Find(&admin, adminID)
	return admin
}

func (db *adminConnection) Users() []entity.User {
	var users []entity.User
	db.connection.Find(&users)
	return users
}
