package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"chat/pkg/config"
	"chat/pkg/utils"
	"chat/services/domain/user/usecase"
	"github.com/gin-gonic/gin"
)

func DeserializeUser(userUseCase usecase.UserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var access_token string
		cookie, err := c.Cookie("access_token")

		authorizationHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)
		fmt.Println("author", fields)
		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			fmt.Println("3")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
		if err != nil {
			fmt.Println("2")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user, err := userUseCase.FindById(fmt.Sprint(sub))
		if err != nil {
			fmt.Println("1")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "The user belonging to this token no logger exists"})
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}
