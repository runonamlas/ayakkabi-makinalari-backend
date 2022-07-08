package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/runonamlas/ayakkabi-makinalari-backend/dto"
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"github.com/runonamlas/ayakkabi-makinalari-backend/helper"
	"github.com/runonamlas/ayakkabi-makinalari-backend/service"
	"net/http"
	"strconv"
)

type AdminController interface {
	Login(ctx *gin.Context)
	Register(tx *gin.Context)
	Update(context *gin.Context)
	Profile(context *gin.Context)
	Users(context *gin.Context)
}

type adminController struct {
	adminService service.AdminService
	jwtService   service.JWTService
}

func NewAdminController(adminService service.AdminService, jwtService service.JWTService) AdminController {
	return &adminController{
		adminService: adminService,
		jwtService:   jwtService,
	}
}

func (c *adminController) Login(ctx *gin.Context) {
	var adminLoginDTO dto.AdminLoginDTO
	errDTO := ctx.ShouldBind(&adminLoginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	adminResult := c.adminService.VerifyCredential(adminLoginDTO.Email)
	if _, ok := adminResult.(entity.Admin); ok {
		passResult := c.adminService.VerifyPassword(adminResult, adminLoginDTO.Password)
		if v, ok := passResult.(entity.Admin); ok {
			generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
			v.Token = generatedToken
			response := helper.BuildResponse(true, "OK!", v)
			ctx.JSON(http.StatusOK, response)
			return
		}
		response := helper.BuildErrorResponse("Please check again your password", "Invalid Credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your email or username", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *adminController) Register(ctx *gin.Context) {
	var adminRegisterDTO dto.AdminRegisterDTO
	errDTO := ctx.ShouldBind(&adminRegisterDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.adminService.IsDuplicateEmail(adminRegisterDTO.Email) || !c.adminService.IsDuplicateUsername(adminRegisterDTO.Username) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email, username and call number", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
	} else {
		createdAdmin := c.adminService.Create(adminRegisterDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdAdmin.ID, 10))
		createdAdmin.Token = token
		response := helper.BuildResponse(true, "OK!", createdAdmin)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *adminController) Update(context *gin.Context) {
	var adminUpdateDTO dto.AdminUpdateDTO
	errDTO := context.ShouldBind(&adminUpdateDTO)
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
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["admin_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	adminUpdateDTO.ID = id
	u := c.adminService.Update(adminUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c *adminController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["admin_id"])
	admin := c.adminService.Profile(id)
	res := helper.BuildResponse(true, "Ok!", admin)
	context.JSON(http.StatusOK, res)
}

func (c *adminController) Users(context *gin.Context) {
	var users = c.adminService.Users()
	res := helper.BuildResponse(true, "OK!", users)
	context.JSON(http.StatusOK, res)
}
