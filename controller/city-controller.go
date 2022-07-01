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

type CityController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type cityController struct {
	cityService service.CityService
	jwtService service.JWTService
}

func NewCityController(cityServ service.CityService, jwtServ service.JWTService) CityController {
	return &cityController{
		cityService: cityServ,
		jwtService: jwtServ,
	}
}

func (c *cityController) All(context *gin.Context) {
	countryID := ""
	countryID = context.Query("country_id")
	if countryID == "" {
		var cities = c.cityService.AllCities()
		res := helper.BuildResponse(true, "OK!", cities)
		context.JSON(http.StatusOK, res)
	}else {
		convertedCountryID, err := strconv.ParseUint(countryID, 10, 64)
		if err == nil {
			var cities = c.cityService.All(convertedCountryID)
			res := helper.BuildResponse(true, "OK!", cities)
			context.JSON(http.StatusOK, res)
		}else{
			res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}
}

func (c *cityController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"),10,64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var city = c.cityService.FindByID(id)
	if (city == entity.City{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	}else {
		res := helper.BuildResponse(true, "OK!", city)
		context.JSON(http.StatusOK, res)
	}
}

func (c *cityController) Insert(context *gin.Context) {
	var cityCreateDTO dto.CityCreateDTO
	errDTO := context.ShouldBind(&cityCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}else {
		result := c.cityService.Insert(cityCreateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *cityController) Update(context *gin.Context) {
	var cityUpdateDTO dto.CityUpdateDTO
	errDTO := context.ShouldBind(&cityUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	result := c.cityService.Update(cityUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	context.JSON(http.StatusOK, response)
}

func (c *cityController) Delete(context *gin.Context) {
	var city entity.City
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No params id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	city.ID = id
	c.cityService.Delete(city)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}