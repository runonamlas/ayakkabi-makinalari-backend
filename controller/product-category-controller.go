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

type ProductCategoryController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type productCategoryController struct {
	productCategoryService service.ProductCategoryService
	jwtService             service.JWTService
}

func NewProductCategoryController(categoryServ service.ProductCategoryService, jwtServ service.JWTService) ProductCategoryController {
	return &productCategoryController{
		productCategoryService: categoryServ,
		jwtService:             jwtServ,
	}
}

func (c *productCategoryController) All(context *gin.Context) {
	var categories = c.productCategoryService.All()
	res := helper.BuildResponse(true, "OK!", categories)
	context.JSON(http.StatusOK, res)
}

func (c *productCategoryController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var category = c.productCategoryService.FindByID(id)
	/*if (category == entity.ProductCategory{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {*/
	res := helper.BuildResponse(true, "OK!", category)
	context.JSON(http.StatusOK, res)
	//}
}

func (c *productCategoryController) Insert(context *gin.Context) {
	var categoryCreateDTO dto.ProductCategoryCreateDTO
	errDTO := context.ShouldBind(&categoryCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.productCategoryService.Insert(categoryCreateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *productCategoryController) Update(context *gin.Context) {
	var categoryUpdateDTO dto.ProductCategoryUpdateDTO
	errDTO := context.ShouldBind(&categoryUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	result := c.productCategoryService.Update(categoryUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	context.JSON(http.StatusOK, response)
}

func (c *productCategoryController) Delete(context *gin.Context) {
	var category entity.ProductCategory
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No params id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	category.ID = id
	c.productCategoryService.Delete(category)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
