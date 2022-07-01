package service

import (
	"fmt"
	"github.com/mashingan/smapping"
	"github.com/my-way-teams/my_way_backend/dto"
	"github.com/my-way-teams/my_way_backend/entity"
	"github.com/my-way-teams/my_way_backend/repository"
	"log"
)

type PlaceService interface {
	Insert(p dto.PlaceCreateDTO) entity.Place
	Update(p dto.PlaceUpdateDTO) entity.Place
	Delete(p entity.Place)
	All(cityID uint64) []entity.Place
	AllPlaces() []entity.Place
	FindByID(placeID uint64) entity.Place
	IsAllowedToEdit(cityID string, placeID uint64) bool
}

type placeService struct {
	placeRepository repository.PlaceRepository
}

func NewPlaceService(placeRepo repository.PlaceRepository) PlaceService {
	return &placeService{
		placeRepository: placeRepo,
	}
}

func (service *placeService) Insert(p dto.PlaceCreateDTO) entity.Place {
	place := entity.Place{}
	err := smapping.FillStruct(&place, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.placeRepository.InsertPlace(place)
	return res
}

func (service *placeService) Update(p dto.PlaceUpdateDTO) entity.Place {
	place := entity.Place{}
	err := smapping.FillStruct(&place, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.placeRepository.UpdatePlace(place)
	return res
}

func (service *placeService) Delete(p entity.Place) {
	service.placeRepository.DeletePlace(p)
}

func (service *placeService) All(cityID uint64) []entity.Place {
	return service.placeRepository.AllPlace(cityID)
}

func (service *placeService) AllPlaces() []entity.Place {
	return service.placeRepository.AllPlaces()
}

func (service *placeService) FindByID(placeID uint64) entity.Place {
	return service.placeRepository.FindPlaceByID(placeID)
}

func (service *placeService) IsAllowedToEdit(cityID string, placeID uint64) bool {
	p := service.placeRepository.FindPlaceByID(placeID)
	id := fmt.Sprintf("%v", p.CityID)
	return cityID == id
}