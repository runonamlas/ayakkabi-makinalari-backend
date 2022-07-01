package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/my-way-teams/my_way_backend/dto"
	"github.com/my-way-teams/my_way_backend/entity"
	"github.com/my-way-teams/my_way_backend/helper"
	"github.com/my-way-teams/my_way_backend/service"
	"net/http"
	"strconv"
)

type PlaceCategoryController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type placeCategoryController struct {
	placeCategoryService service.PlaceCategoryService
	jwtService service.JWTService
}

func NewPlaceCategoryController(categoryServ service.PlaceCategoryService, jwtServ service.JWTService) PlaceCategoryController {
	return &placeCategoryController{
		placeCategoryService: categoryServ,
		jwtService: jwtServ,
	}
}

func (c *placeCategoryController) All(context *gin.Context) {
	var categories = c.placeCategoryService.All()
	res := helper.BuildResponse(true, "OK!", categories)
	context.JSON(http.StatusOK, res)
}

func (c *placeCategoryController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"),0,0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var category = c.placeCategoryService.FindByID(id)
	if (category == entity.PlaceCategory{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	}else {
		res := helper.BuildResponse(true, "OK!", category)
		context.JSON(http.StatusOK, res)
	}
}

func (c *placeCategoryController) Insert(context *gin.Context) {
	var categoryCreateDTO dto.PlaceCategoryCreateDTO
	errDTO := context.ShouldBind(&categoryCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}else {
		result := c.placeCategoryService.Insert(categoryCreateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *placeCategoryController) Update(context *gin.Context) {
	var categoryUpdateDTO dto.PlaceCategoryUpdateDTO
	errDTO := context.ShouldBind(&categoryUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	result := c.placeCategoryService.Update(categoryUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	context.JSON(http.StatusOK, response)
}

func (c *placeCategoryController) Delete(context *gin.Context) {
	var category entity.PlaceCategory
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No params id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	category.ID = id
	c.placeCategoryService.Delete(category)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}