package http

import (
	"chat/domain"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func NewLoginHandler(r *gin.Engine, login domain.LoginUsecase) {
	handler := LoginHanlder{login}

	r.GET("/login", handler.CreateJwtUser)
}

func (login *LoginHanlder) CreateJwtUser(c *gin.Context) {
	var user domain.Login

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
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "user is not exists",
		})
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &domain.Claims{
		ID:    res.ID.String(),
		Email: res.Email,
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
		"status":  http.StatusOK,
		"id":      res.ID,
		"token":   tokenString,
		"expires": expirationTime.Unix(),
	})
}
