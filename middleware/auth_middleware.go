package middleware

import (
	"fmt"
	"net/http"

	"github.com/YashBubna/jwt_authentication/helpers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("No Authorization")})
			c.Abort()
			return
		}
		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.Render(http.StatusInternalServerError, render.JSON{Data: any(gin.H{"error": err})})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("user_type", claims.User_type)
		c.Set("user_id", claims.User_id)
		c.Next()
	}
}
