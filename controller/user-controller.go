package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/runonamlas/ayakkabi-makinalari-backend/dto"
	"github.com/runonamlas/ayakkabi-makinalari-backend/helper"
	"github.com/runonamlas/ayakkabi-makinalari-backend/service"
	"net/http"
	"strconv"
)

type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
	GetProducts(context *gin.Context)
	//AddFavourite(context *gin.Context)
	//DeleteFavourite(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) GetProducts(context *gin.Context) {
	countryID := ""
	countryID = context.Query("country_id")
	if countryID != "" {
		authHeader := context.GetHeader("Authorization")
		token, err := c.jwtService.ValidateToken(authHeader)
		if err != nil {
			panic(err.Error())
		}
		claims := token.Claims.(jwt.MapClaims)
		id := fmt.Sprintf("%v", claims["user_id"])
		userFavorites := c.userService.GetProducts(id)
		res := helper.BuildResponse(true, "Ok!", userFavorites)
		context.JSON(http.StatusOK, res)
	} else {
		res := helper.BuildErrorResponse("No param id was found", "Please country id", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
}

/*func (c *userController) AddFavourite(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	placeID, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No params id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	userFavorites := c.userService.AddFavourite(id, placeID)
	res := helper.BuildResponse(true, "Ok!", userFavorites)
	context.JSON(http.StatusOK, res)
}

func (c *userController) DeleteFavourite(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	placeID, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No params id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	deleteFavourite := c.userService.DeleteFavourite(id, placeID)
	res := helper.BuildResponse(true, "Delete!", deleteFavourite)
	context.JSON(http.StatusOK, res)
}*/

func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "Ok!", user)
	context.JSON(http.StatusOK, res)
}
