package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/runonamlas/ayakkabi-makinalari-backend/dto"
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"github.com/runonamlas/ayakkabi-makinalari-backend/helper"
	"github.com/runonamlas/ayakkabi-makinalari-backend/service"
	"net/http"
	"strconv"
)

type ProductController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type productController struct {
	productService service.ProductService
	jwtService     service.JWTService
}

func NewProductController(productServ service.ProductService, jwtServ service.JWTService) ProductController {
	return &productController{
		productService: productServ,
		jwtService:     jwtServ,
	}
}

func (p *productController) All(context *gin.Context) {
	cityID := ""
	cityID = context.Query("city_id")
	if cityID == "" {
		var products = p.productService.AllProducts()
		res := helper.BuildResponse(true, "OK!", products)
		context.JSON(http.StatusOK, res)
	} else {
		convertedCityID, err := strconv.ParseUint(cityID, 10, 64)
		if err == nil {
			var products = p.productService.All(convertedCityID)
			res := helper.BuildResponse(true, "OK!", products)
			context.JSON(http.StatusOK, res)
		} else {
			res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}
}

func (p *productController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var product = p.productService.FindByID(id)
	//if (place == entity.Place{}) {
	//	res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
	//	context.JSON(http.StatusNotFound, res)
	//}else {
	res := helper.BuildResponse(true, "OK!", product)
	context.JSON(http.StatusOK, res)
}

func (p *productController) Insert(context *gin.Context) {
	var productCreateDTO dto.ProductCreateDTO
	errDTO := context.ShouldBind(&productCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := p.productService.Insert(productCreateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (p *productController) Update(context *gin.Context) {
	var productUpdateDTO dto.ProductUpdateDTO
	errDTO := context.ShouldBind(&productUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	result := p.productService.Update(productUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	context.JSON(http.StatusOK, response)
}

func (p *productController) Delete(context *gin.Context) {
	var product entity.Product
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No params id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	product.ID = id
	p.productService.Delete(product)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
