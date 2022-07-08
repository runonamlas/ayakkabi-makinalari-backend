package service

import (
	"github.com/mashingan/smapping"
	"github.com/runonamlas/ayakkabi-makinalari-backend/dto"
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"github.com/runonamlas/ayakkabi-makinalari-backend/repository"
	"log"
)

type ProductCategoryService interface {
	Insert(c dto.ProductCategoryCreateDTO) entity.ProductCategory
	Update(c dto.ProductCategoryUpdateDTO) entity.ProductCategory
	Delete(c entity.ProductCategory)
	All() []entity.ProductCategory
	FindByID(categoryID uint64) entity.ProductCategory
}

type productCategoryService struct {
	productCategoryRepository repository.ProductCategoryRepository
}

func NewProductCategoryService(categoryRepo repository.ProductCategoryRepository) ProductCategoryService {
	return &productCategoryService{
		productCategoryRepository: categoryRepo,
	}
}

func (service *productCategoryService) Insert(c dto.ProductCategoryCreateDTO) entity.ProductCategory {
	category := entity.ProductCategory{}
	err := smapping.FillStruct(&category, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.productCategoryRepository.InsertCategory(category)
	return res
}

func (service *productCategoryService) Update(c dto.ProductCategoryUpdateDTO) entity.ProductCategory {
	category := entity.ProductCategory{}
	err := smapping.FillStruct(&category, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.productCategoryRepository.UpdateCategory(category)
	return res
}

func (service *productCategoryService) Delete(c entity.ProductCategory) {
	service.productCategoryRepository.DeleteCategory(c)
}

func (service *productCategoryService) All() []entity.ProductCategory {
	return service.productCategoryRepository.AllCategory()
}

func (service *productCategoryService) FindByID(categoryID uint64) entity.ProductCategory {
	return service.productCategoryRepository.FindCategoryByID(categoryID)
}
