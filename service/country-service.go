package service

import (
	"github.com/mashingan/smapping"
	"github.com/my-way-teams/my_way_backend/dto"
	"github.com/my-way-teams/my_way_backend/entity"
	"github.com/my-way-teams/my_way_backend/repository"
	"log"
)

type CountryService interface {
	Insert(c dto.CountryCreateDTO) entity.Country
	Update(c dto.CountryUpdateDTO) entity.Country
	Delete(c entity.Country)
	All() []entity.Country
	FindByID(countryID uint64) entity.Country
}

type countryService struct {
	countryRepository repository.CountryRepository
}

func NewCountryService(countryRepo repository.CountryRepository) CountryService {
	return &countryService{
		countryRepository: countryRepo,
	}
}

func (service *countryService) Insert(c dto.CountryCreateDTO) entity.Country {
	country := entity.Country{}
	err := smapping.FillStruct(&country, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.countryRepository.InsertCountry(country)
	return res
}

func (service *countryService) Update(c dto.CountryUpdateDTO) entity.Country {
	country := entity.Country{}
	err := smapping.FillStruct(&country, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.countryRepository.UpdateCountry(country)
	return res
}

func (service *countryService) Delete(c entity.Country) {
	service.countryRepository.DeleteCountry(c)
}

func (service *countryService) All() []entity.Country {
	return service.countryRepository.AllCountry()
}

func (service *countryService) FindByID(countryID uint64) entity.Country {
	return service.countryRepository.FindCountryByID(countryID)
}