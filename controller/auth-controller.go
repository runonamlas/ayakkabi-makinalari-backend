package controller

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/runonamlas/ayakkabi-makinalari-backend/dto"
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"github.com/runonamlas/ayakkabi-makinalari-backend/helper"
	"github.com/runonamlas/ayakkabi-makinalari-backend/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Forget(ctx *gin.Context)
	Change(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email)
	if _, ok := authResult.(entity.User); ok {
		passResult := c.authService.VerifyPassword(authResult, loginDTO.Password)
		if v, ok := passResult.(entity.User); ok {
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

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) || !c.authService.IsDuplicateUsername(registerDTO.Username) || !c.authService.IsDuplicateCallNumber(registerDTO.CallNumber) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email, username and call number", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *authController) Forget(ctx *gin.Context) {
	var forgetDTO dto.ForgetDTO
	errDTO := ctx.ShouldBind(&forgetDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(forgetDTO.Email) {
		authResult := c.authService.VerifyCredential(forgetDTO.Email)
		if v, ok := authResult.(entity.User); ok {
			generatedToken := c.jwtService.GenerateTokenForget(strconv.FormatUint(v.ID, 10))
			v.Token = generatedToken
			from := "ayakkabimakineleri@gmail.com"
			password := "doblaqyyqlvrboeo"
			to := []string{v.Email}
			smtpHost := "smtp.gmail.com"
			smtpPort := "587"
			message := []byte("To: " + v.Email + "\r\n" +
				"Subject: Şifremi Unuttum\r\n" +
				"\r\n" + "Eğer şifre sıfırlama isteğini siz yolladıysanız lütfen aşağıdaki linke tıklayınız. Eğer siz şifrenizi sıfırlamak istemediyseniz bu maili görmezden gelini  link : http:localhost:3000/sifreyi-yenile/" + generatedToken + " \r\n")
			auth := smtp.PlainAuth("", from, password, smtpHost)
			err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
			if err != nil {
				log.Fatal(err)
				response := helper.BuildErrorResponse("Failed to process", "Not success", helper.EmptyObj{})
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			}
			response := helper.BuildResponse(true, "ok", helper.EmptyObj{})
			ctx.JSON(http.StatusOK, response)
			return
		}
		response := helper.BuildErrorResponse("Kullanıcı adı şifre eşleşmiyor", "Invalid Credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	} else {
		response := helper.BuildErrorResponse("Failed to process request", "Not found email", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
	}
}

func (c *authController) Change(ctx *gin.Context) {
	var changePasswordDTO dto.ChangePasswordDTO
	errDTO := ctx.ShouldBind(&changePasswordDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		res := helper.BuildErrorResponse("Failed to process request", errToken.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	changePasswordDTO.ID = id
	u := c.authService.ChangePassword(changePasswordDTO)
	res := helper.BuildResponse(true, "OK!", u)
	ctx.JSON(http.StatusOK, res)
}
