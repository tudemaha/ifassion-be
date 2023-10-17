package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tudemaha/ifassion-be/internal/global/responses"
)

func APIKeyValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response responses.Response

		auth := c.GetHeader("Authorization")
		fmt.Println(auth)
		if auth == "" {
			response.DefaultUnauthorized()
			c.AbortWithStatusJSON(response.Code, response)
			return
		}

		api_key := strings.Split(auth, " ")
		if api_key[0] != "Bearer" {
			response.DefaultUnauthorized()
			c.AbortWithStatusJSON(response.Code, response)
			return
		}

		if api_key[1] != os.Getenv("API_KEY") {
			response.DefaultUnauthorized()
			c.AbortWithStatusJSON(response.Code, response)
			return
		}

		c.Next()
	}
}
