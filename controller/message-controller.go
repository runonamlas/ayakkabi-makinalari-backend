package controller

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/runonamlas/ayakkabi-makinalari-backend/dto"
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"github.com/runonamlas/ayakkabi-makinalari-backend/helper"
	"github.com/runonamlas/ayakkabi-makinalari-backend/service"
)

type MessageController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type messageController struct {
	messageService service.MessageService
	jwtService     service.JWTService
}

func NewMessageController(messageServ service.MessageService, jwtServ service.JWTService) MessageController {
	return &messageController{
		messageService: messageServ,
		jwtService:     jwtServ,
	}
}

func (p *messageController) All(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := p.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := claims["user_id"]
	ownerID, err := strconv.ParseInt(id.(string), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	userID, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var messages = p.messageService.All(userID, uint64(ownerID))
	res := helper.BuildResponse(true, "OK!", messages)
	context.JSON(http.StatusOK, res)

}

func (p *messageController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var message = p.messageService.FindByID(id)
	//if (place == entity.Place{}) {
	//	res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
	//	context.JSON(http.StatusNotFound, res)
	//}else {
	res := helper.BuildResponse(true, "OK!", message)
	context.JSON(http.StatusOK, res)
}

func (p *messageController) Insert(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := p.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := claims["user_id"]
	n, err := strconv.ParseInt(id.(string), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	var messageCreateDTO dto.MessageCreateDTO
	errDTO := context.ShouldBind(&messageCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		messageCreateDTO.OwnerID = uint64(n)
		result := p.messageService.Insert(messageCreateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (p *messageController) Update(context *gin.Context) {
	var messageUpdateDTO dto.MessageUpdateDTO
	errDTO := context.ShouldBind(&messageUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	result := p.messageService.Update(messageUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	context.JSON(http.StatusOK, response)
}

func (p *messageController) Delete(context *gin.Context) {
	var message entity.Message
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No params id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	message.ID = id
	p.messageService.Delete(message)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
