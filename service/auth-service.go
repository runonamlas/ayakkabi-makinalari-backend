package service

import (
	"github.com/mashingan/smapping"
	"github.com/runonamlas/ayakkabi-makinalari-backend/dto"
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"github.com/runonamlas/ayakkabi-makinalari-backend/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService interface {
	VerifyCredential(email string) interface{}
	VerifyPassword(res interface{}, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
	FindByUsername(username string) entity.User
	IsDuplicateUsername(username string) bool
	IsDuplicateCallNumber(callNumber string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email string) interface{} {
	res := service.userRepository.VerifyCredential(email)
	if v, ok := res.(entity.User); ok {
		if v.Email == email || v.CallNumber == email {
			return res
		}
		return false
	}
	return false
}

func (service *authService) VerifyPassword(res interface{}, password string) interface{} {
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparedPassword(v.Password, []byte(password))
		if comparedPassword {
			return res
		}
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (service *authService) FindByUsername(username string) entity.User {
	return service.userRepository.FindByUsername(username)
}

func (service *authService) IsDuplicateUsername(username string) bool {
	res := service.userRepository.IsDuplicateUsername(username)
	return !(res.Error == nil)
}

func (service *authService) IsDuplicateCallNumber(callNumber string) bool {
	res := service.userRepository.IsDuplicateCallNumber(callNumber)
	return !(res.Error == nil)
}

func comparedPassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
