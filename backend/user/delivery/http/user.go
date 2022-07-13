package http

import (
	"chat/domain"
	"chat/jwt"
	"net/http"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUse domain.UserUsecase
}

func NewUserHandler(r *gin.Engine, userUsecase domain.UserUsecase) {
	userHandle := &userHandler{userUsecase}
	r.GET("/users", userHandle.GetAll)
	r.GET("/users/:id", userHandle.Welcome)
}

func (u *userHandler) GetAll(c *gin.Context) {
	users := u.userUse.GetAllUser()

	if users == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      "failed",
			"status_code": http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      "successful",
		"status_code": http.StatusOK,
		"data":        users,
	})
}
func (u *userHandler) Welcome(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Cookie is not available",
		})
		return
	}
	token := cookie.Value
	// Decode token
	claims, err := jwt.Decode(token)

	id := c.Param("id")
	if id == claims.ID {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"Welcome": claims.Email,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID is not sample",
		})
	}

}
func (u *userHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
