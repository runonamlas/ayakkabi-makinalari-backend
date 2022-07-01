package service

import (
	"fmt"
	"github.com/mashingan/smapping"
	"github.com/my-way-teams/my_way_backend/dto"
	"github.com/my-way-teams/my_way_backend/entity"
	"github.com/my-way-teams/my_way_backend/repository"
	"log"
)

type CityService interface {
	Insert(c dto.CityCreateDTO) entity.City
	Update(c dto.CityUpdateDTO) entity.City
	Delete(c entity.City)
	All(countryID uint64) []entity.City
	AllCities() []entity.City
	FindByID(cityID uint64) entity.City
	IsAllowedToEdit(countryID string, cityID uint64) bool
}

type cityService struct {
	cityRepository repository.CityRepository
}

func NewCityService(cityRepo repository.CityRepository) CityService {
	return &cityService{
		cityRepository: cityRepo,
	}
}

func (service *cityService) Insert(c dto.CityCreateDTO) entity.City {
	city := entity.City{}
	err := smapping.FillStruct(&city, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.cityRepository.InsertCity(city)
	return res
}

func (service *cityService) Update(c dto.CityUpdateDTO) entity.City {
	city := entity.City{}
	err := smapping.FillStruct(&city, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.cityRepository.UpdateCity(city)
	return res
}

func (service *cityService) Delete(c entity.City) {
	service.cityRepository.DeleteCity(c)
}

func (service *cityService) All(countryID uint64) []entity.City {
	return service.cityRepository.AllCity(countryID)
}

func (service *cityService) AllCities() []entity.City {
	return service.cityRepository.AllCities()
}

func (service *cityService) FindByID(cityID uint64) entity.City {
	return service.cityRepository.FindCityByID(cityID)
}

func (service *cityService) IsAllowedToEdit(countryID string, cityID uint64) bool {
	c := service.cityRepository.FindCityByID(cityID)
	id := fmt.Sprintf("%v", c.CountryID)
	return countryID == id
}