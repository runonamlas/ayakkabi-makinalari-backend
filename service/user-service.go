package service

import (
	"github.com/mashingan/smapping"
	"github.com/runonamlas/ayakkabi-makinalari-backend/dto"
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"github.com/runonamlas/ayakkabi-makinalari-backend/repository"
	"log"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
	GetProducts(userID string) []entity.Product
	//AddFavourite(userID string, placeID uint64) entity.
	//DeleteFavourite(userID string, placeID uint64) bool
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

func (service *userService) GetProducts(userID string) []entity.Product {
	return service.userRepository.GetProducts(userID)
}

/*func (service *userService) AddFavourite(userID string, placeID uint64) entity.Place {
	return service.userRepository.AddFavourite(userID, placeID)
}

func (service *userService) DeleteFavourite(userID string, placeID uint64) bool {
	return service.userRepository.DeleteFavourite(userID, placeID)
}
*/
