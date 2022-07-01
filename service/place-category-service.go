package service

import (
	"github.com/mashingan/smapping"
	"github.com/my-way-teams/my_way_backend/dto"
	"github.com/my-way-teams/my_way_backend/entity"
	"github.com/my-way-teams/my_way_backend/repository"
	"log"
)

type PlaceCategoryService interface {
	Insert(c dto.PlaceCategoryCreateDTO) entity.PlaceCategory
	Update(c dto.PlaceCategoryUpdateDTO) entity.PlaceCategory
	Delete(c entity.PlaceCategory)
	All() []entity.PlaceCategory
	FindByID(categoryID uint64) entity.PlaceCategory
}

type placeCategoryService struct {
	placeCategoryRepository repository.PlaceCategoryRepository
}

func NewPlaceCategoryService(categoryRepo repository.PlaceCategoryRepository) PlaceCategoryService {
	return &placeCategoryService{
		placeCategoryRepository: categoryRepo,
	}
}

func (service *placeCategoryService) Insert(c dto.PlaceCategoryCreateDTO) entity.PlaceCategory {
	category := entity.PlaceCategory{}
	err := smapping.FillStruct(&category, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.placeCategoryRepository.InsertCategory(category)
	return res
}

func (service *placeCategoryService) Update(c dto.PlaceCategoryUpdateDTO) entity.PlaceCategory {
	category := entity.PlaceCategory{}
	err := smapping.FillStruct(&category, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.placeCategoryRepository.UpdateCategory(category)
	return res
}

func (service *placeCategoryService) Delete(c entity.PlaceCategory) {
	service.placeCategoryRepository.DeleteCategory(c)
}

func (service *placeCategoryService) All() []entity.PlaceCategory {
	return service.placeCategoryRepository.AllCategory()
}

func (service *placeCategoryService) FindByID(categoryID uint64) entity.PlaceCategory {
	return service.placeCategoryRepository.FindCategoryByID(categoryID)
}
