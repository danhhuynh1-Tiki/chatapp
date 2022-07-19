package http

import (
	"chat/domain"
	"chat/jwt"
	"chat/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginHanlder struct {
	loginUse domain.LoginUsecase
}

// type Claims struct {
// 	ID    string `json:"id"`
// 	Email string `json:"email"`
// 	jwt.StandardClaims
// }

func NewLoginHandler(r *gin.RouterGroup, login domain.LoginUsecase) {
	handler := LoginHanlder{login}

	r.GET("/login", handler.CreateJwtUser)
}

func (login *LoginHanlder) CreateJwtUser(c *gin.Context) {
	var user domain.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "require data",
		})
		return
	}
	res, err := login.loginUse.GetUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorNotFound(err))
		return
	}
	// update status and time request
	t := time.Now()
	err = login.loginUse.UpdateStatusUser(res.ID, t, 1)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "Cannot update status",
		})
	}
	tokenString, expirationTime, err := jwt.Encode(*res)

	// set cookie for jwt
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"id":      res.ID,
		"token":   tokenString,
		"expires": expirationTime.Unix(),
	})
}
