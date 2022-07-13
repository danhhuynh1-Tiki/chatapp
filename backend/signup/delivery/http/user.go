package controllers

import (
	"chat/domain"
	"time"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	responses "chat/responses"
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
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &domain.Claims{
		ID:    user.ID.String(),
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(domain.JwtKey)
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
