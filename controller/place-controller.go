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

type PlaceController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type placeController struct {
	placeService service.PlaceService
	jwtService service.JWTService
}

func NewPlaceController(placeServ service.PlaceService, jwtServ service.JWTService) PlaceController {
	return &placeController{
		placeService: placeServ,
		jwtService: jwtServ,
	}
}

func (p *placeController) All(context *gin.Context) {
	cityID := ""
	cityID = context.Query("city_id")
	if cityID == "" {
		var places = p.placeService.AllPlaces()
		res := helper.BuildResponse(true, "OK!", places)
		context.JSON(http.StatusOK, res)
	}else {
		convertedCityID, err := strconv.ParseUint(cityID, 10, 64)
		if err == nil {
			var places = p.placeService.All(convertedCityID)
			res := helper.BuildResponse(true, "OK!", places)
			context.JSON(http.StatusOK, res)
		}else{
			res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}
}

func (p *placeController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"),10,64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var place = p.placeService.FindByID(id)
	//if (place == entity.Place{}) {
	//	res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
	//	context.JSON(http.StatusNotFound, res)
	//}else {
	res := helper.BuildResponse(true, "OK!", place)
	context.JSON(http.StatusOK, res)
}

func (p *placeController) Insert(context *gin.Context) {
	var placeCreateDTO dto.PlaceCreateDTO
	errDTO := context.ShouldBind(&placeCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}else {
		result := p.placeService.Insert(placeCreateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (p *placeController) Update(context *gin.Context) {
	var placeUpdateDTO dto.PlaceUpdateDTO
	errDTO := context.ShouldBind(&placeUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	result := p.placeService.Update(placeUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	context.JSON(http.StatusOK, response)
}

func (p *placeController) Delete(context *gin.Context) {
	var place entity.Place
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No params id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	place.ID = id
	p.placeService.Delete(place)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}