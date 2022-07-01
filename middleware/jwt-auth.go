package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/my-way-teams/my_way_backend/helper"
	"github.com/my-way-teams/my_way_backend/service"
	"log"
	"net/http"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc  {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if !token.Valid {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}