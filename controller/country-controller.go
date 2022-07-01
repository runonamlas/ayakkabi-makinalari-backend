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

type CountryController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type countryController struct {
	countryService service.CountryService
	jwtService service.JWTService
}

func NewCountryController(countryServ service.CountryService, jwtServ service.JWTService) CountryController {
	return &countryController{
		countryService: countryServ,
		jwtService: jwtServ,
	}
}

func (c *countryController) All(context *gin.Context) {
	var countries = c.countryService.All()
	res := helper.BuildResponse(true, "OK!", countries)
	context.JSON(http.StatusOK, res)
}

func (c *countryController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"),0,0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var country = c.countryService.FindByID(id)
	if (country == entity.Country{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	}else {
		res := helper.BuildResponse(true, "OK!", country)
		context.JSON(http.StatusOK, res)
	}
}

func (c *countryController) Insert(context *gin.Context) {
	var countryCreateDTO dto.CountryCreateDTO
	errDTO := context.ShouldBind(&countryCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}else {
		result := c.countryService.Insert(countryCreateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *countryController) Update(context *gin.Context) {
	var countryUpdateDTO dto.CountryUpdateDTO
	errDTO := context.ShouldBind(&countryUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	result := c.countryService.Update(countryUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	context.JSON(http.StatusOK, response)
}

func (c *countryController) Delete(context *gin.Context) {
	var country entity.Country
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No params id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	country.ID = id
	c.countryService.Delete(country)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}