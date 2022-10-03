package service

import (
	"github.com/mashingan/smapping"
	"github.com/runonamlas/ayakkabi-makinalari-backend/dto"
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"github.com/runonamlas/ayakkabi-makinalari-backend/repository"
	"log"
)

type ProductService interface {
	Insert(p dto.ProductCreateDTO) entity.Product
	Update(p dto.ProductUpdateDTO) entity.Product
	Delete(p entity.Product)
	Sold(p entity.Product)
	All(cityID uint64) []entity.Product
	AllProducts() []entity.Product
	FindByID(productID uint64) entity.Product
	FindByCategory(categoryID uint64) []entity.Product
	IsAllowedToEdit(cityID string, productID uint64) bool
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepo,
	}
}

func (service *productService) Insert(p dto.ProductCreateDTO) entity.Product {
	product := entity.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.productRepository.InsertProduct(product)
	return res
}

func (service *productService) Update(p dto.ProductUpdateDTO) entity.Product {
	product := entity.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.productRepository.UpdateProduct(product)
	return res
}

func (service *productService) Delete(p entity.Product) {
	service.productRepository.DeleteProduct(p)
}

func (service *productService) Sold(p entity.Product) {
	service.productRepository.SoldProduct(p)
}

func (service *productService) All(cityID uint64) []entity.Product {
	return service.productRepository.AllProduct(cityID)
}

func (service *productService) AllProducts() []entity.Product {
	return service.productRepository.AllProducts()
}

func (service *productService) FindByID(productID uint64) entity.Product {
	return service.productRepository.FindProductByID(productID)
}

func (service *productService) FindByCategory(categoryID uint64) []entity.Product {
	return service.productRepository.FindProductByCategory(categoryID)
}

func (service *productService) IsAllowedToEdit(cityID string, productID uint64) bool {
	//p := service.productRepository.FindProductByID(productID)
	//id := fmt.Sprintf("%v", p.CityID)
	//return cityID == id
	return false
}
