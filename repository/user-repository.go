package repository

import (
	"log"

	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	Statistic(userID string) entity.User
	GetProducts(userID string) []entity.Product
	GetMessages(userID string) []entity.Message
	ChangePassword(user entity.User) bool
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
	var tempUser entity.User
	db.connection.Find(&tempUser, user.ID)
	if user.Username != "" {
		tempUser.Username = user.Username
	}
	if user.Address != "" {
		tempUser.Address = user.Address
	}
	db.connection.Save(&tempUser)
	return tempUser
}

func (db *userConnection) ChangePassword(user entity.User) bool {
	var tempUser entity.User
	db.connection.Find(&tempUser, user.ID)
	tempUser.Password = hashAndSalt([]byte(user.Password))
	res := db.connection.Save(&tempUser)
	return res.Error == nil
}

func (db *userConnection) VerifyCredential(email string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	resUsername := db.connection.Where("call_number = ?", email).Take(&user)
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
	return db.connection.Where("call_number = ?", callNumber).Take(&user)
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
	db.connection.Preload("Products").Where("username = ?", userID).Take(&user)
	user.ClickProfile = user.ClickProfile + 1
	db.connection.Save(&user)
	return user
}

func (db *userConnection) Statistic(userID string) entity.User {
	var user entity.User
	db.connection.Preload("Products").Find(&user, userID)
	return user
}

func (db *userConnection) GetProducts(userID string) []entity.Product {
	var user entity.User
	var products []entity.Product
	db.connection.Preload("Products.Category").Find(&user, userID)
	for _, v := range user.Products {
		products = append(products, *v)
	}
	return products
}

func (db *userConnection) GetMessages(userID string) []entity.Message {
	var user entity.User
	var messages []entity.Message
	db.connection.Preload("Messages.Owner").Preload("Messages.Product").Preload("Messages.User").Find(&user, userID)
	for _, v := range user.Messages {
		var a = false
		for i, message := range messages {
			if v.OwnerID == message.OwnerID {
				messages[i] = *v
				a = true
				break
			}
		}
		if !a {
			messages = append(messages, *v)
		}
	}
	return messages
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
