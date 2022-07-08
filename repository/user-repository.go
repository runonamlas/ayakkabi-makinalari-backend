package repository

import (
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	IsDuplicateUsername(username string) (tx *gorm.DB)
	IsDuplicateCallNumber(callNumber string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	FindByUsername(username string) entity.User
	ProfileUser(userID string) entity.User
	GetProducts(userID string) []entity.Product
	//AddFavourite(userID string, placeID uint64) entity.Place
	//DeleteFavourite(userID string, placeID uint64) bool
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	resUsername := db.connection.Where("username = ?", email).Take(&user)
	if res.Error == nil || resUsername.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) IsDuplicateUsername(username string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("username = ?", username).Take(&user)
}

func (db *userConnection) IsDuplicateCallNumber(callNumber string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("callNumber = ?", callNumber).Take(&user)
}

func (db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) FindByUsername(username string) entity.User {
	var user entity.User
	db.connection.Where("username = ?", username).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) entity.User {
	var user entity.User
	db.connection.Find(&user, userID)
	return user
}

func (db *userConnection) GetProducts(userID string) []entity.Product {
	var user entity.User
	var places []entity.Product
	println("-------")
	db.connection.Preload("Products").Find(&user, userID)
	/*for _, v := range user.Products {
		if v.City.CountryID == countryID {
			places = append(places, v)
		}
	}*/
	return places
}

/*func (db *userConnection) AddFavourite(userID string, placeID uint64) entity.Place {
	var user entity.User
	var place entity.Place
	db.connection.Preload("Places").Find(&user, userID)
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Find(&place, placeID)
	user.Places = append(user.Places, place)
	db.connection.Save(&user)
	return place
}

func (db *userConnection) DeleteFavourite(userID string, placeID uint64) bool {
	var user entity.User
	var place entity.Place
	db.connection.Preload("Places").Find(&user, userID)
	db.connection.Preload("Routes").Preload("City.Country").Preload("Category").Find(&place, placeID)

	db.connection.Model(&user).Association("Places").Delete(place)
	//error handle yapilacak
	return true
}*/

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash password")
	}
	return string(hash)
}
