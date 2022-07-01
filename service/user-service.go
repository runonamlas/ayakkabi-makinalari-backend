package service

import (
	"github.com/mashingan/smapping"
	"github.com/my-way-teams/my_way_backend/dto"
	"github.com/my-way-teams/my_way_backend/entity"
	"github.com/my-way-teams/my_way_backend/repository"
	"log"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
	GetFavourites(userID string, countryID uint64) []entity.Place
	AddFavourite(userID string, placeID uint64) entity.Place
	DeleteFavourite(userID string, placeID uint64) bool
	GetSaved(userID string, countryID uint64) []entity.Route
	AddSaved(userID string, routeID uint64) entity.Route
	DeleteSaved(userID string, routeID uint64) bool
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}

func (service *userService) GetFavourites(userID string, countryID uint64) []entity.Place {
	return service.userRepository.GetFavourites(userID, countryID)
}

func (service *userService) AddFavourite(userID string, placeID uint64) entity.Place {
	return service.userRepository.AddFavourite(userID, placeID)
}

func (service *userService) DeleteFavourite(userID string, placeID uint64) bool {
	return service.userRepository.DeleteFavourite(userID, placeID)
}

func (service *userService) GetSaved(userID string, countryID uint64) []entity.Route {
	return service.userRepository.GetSaved(userID, countryID)
}

func (service *userService) AddSaved(userID string, routeID uint64) entity.Route {
	return service.userRepository.AddSaved(userID, routeID)
}

func (service *userService) DeleteSaved(userID string, routeID uint64) bool {
	return service.userRepository.DeleteSaved(userID, routeID)
}