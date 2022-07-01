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

type RouteController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type routeController struct {
	routeService service.RouteService
	jwtService service.JWTService
}

func NewRouteController(routeServ service.RouteService, jwtServ service.JWTService) RouteController {
	return &routeController{
		routeService: routeServ,
		jwtService: jwtServ,
	}
}

func (r *routeController) All(context *gin.Context) {
	cityID := context.Query("city_id")

	if cityID == "" {
		var routes = r.routeService.AllRoutes()
		res := helper.BuildResponse(true, "OK!", routes)
		context.JSON(http.StatusOK, res)
	}else {
		convertedCityID, err := strconv.ParseUint(cityID, 10, 64)
		if err == nil {
			var routes = r.routeService.All(convertedCityID)
			res := helper.BuildResponse(true, "OK!", routes)
			context.JSON(http.StatusOK, res)
		}else{
			res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}
}

func (r *routeController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"),10,64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var route = r.routeService.FindByID(id)
	res := helper.BuildResponse(true, "OK!", route)
	context.JSON(http.StatusOK, res)
}

func (r *routeController) Insert(context *gin.Context) {
	var routeCreateDTO dto.RouteCreateDTO
	errDTO := context.ShouldBind(&routeCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}else {
		result := r.routeService.Insert(routeCreateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (r *routeController) Update(context *gin.Context) {
	var routeUpdateDTO dto.RouteUpdateDTO
	errDTO := context.ShouldBind(&routeUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	result := r.routeService.Update(routeUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	context.JSON(http.StatusOK, response)
}

func (r *routeController) Delete(context *gin.Context) {
	var route entity.Route
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No params id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	route.ID = id
	r.routeService.Delete(route)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}