package controllers

import (
	"chat/domain"

	"net/http"

	jwt "chat/jwt"
	responses "chat/responses"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase domain.SingupUsecase
}

func NewSignupHandler(r *gin.Engine, userUsecase domain.SingupUsecase) {
	userHandler := &UserController{userUsecase}
	r.POST("/signup", userHandler.Create)
}

func (userCtl UserController) Create(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorInvalidData(err))
		return
	}
	_, err := userCtl.userUsecase.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorInvalidData(err))
		return
	}
	// get token string
	tokenString, expirationTime, err := jwt.Encode(user)
	// set cookie for jwt
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	c.JSON(http.StatusOK, gin.H{
		"status":      "successful",
		"status_code": http.StatusOK,
		"id":          user.ID,
		"token":       tokenString,
		"expires":     expirationTime.Unix(),
	})

}
