package http

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	config "chat/pkg/config"
	filter "chat/pkg/filter"
	utils "chat/pkg/utils"
	authUseCase "chat/services/domain/auth/usecase"
	userUseCase "chat/services/domain/user/usecase"
	"chat/services/models"
)

type AuthHandler struct {
	authUsecase authUseCase.AuthUseCase
	userUsecase userUseCase.UserUseCase
}

func NewAuthHandler(authUsecase authUseCase.AuthUseCase, userUsecase userUseCase.UserUseCase) AuthHandler {
	return AuthHandler{
		authUsecase: authUsecase,
		userUsecase: userUsecase,
	}
}

func (handler *AuthHandler) SignUpUser(c *gin.Context) {
	var user *models.SignUpInput

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	// set status user
	user.Status = 1
	//if user.Password != user.PasswordConfirm {
	//	c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "passwords do not match"})
	//	return
	//}

	newUser, err := handler.authUsecase.SignUp(user)
	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			c.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{
		"user": filter.FilteredResponse(newUser),
	}})
}

func (handler *AuthHandler) SignInUser(c *gin.Context) {
	var credentials *models.SignInInput
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	user, err := handler.userUsecase.FindByEmail(credentials.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := utils.VerifyPassword(user.Password, credentials.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}
	// Update status
	query := user.ID
	err = handler.userUsecase.UpdateStatus(query, 1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "meesage": "Cannot update status"})
		return
	}
	config, _ := config.LoadConfig(".")
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "access_token",
		Value:   access_token,
		Expires: time.Now().Add(5 * time.Minute),
	})
	c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "", false, true)
	c.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "", false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "", false, false)

	c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}

func (handler *AuthHandler) RefreshAccessToken(c *gin.Context) {
	message := "could not refresh access token"
	cookie, err := c.Cookie("refresh_token")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
		return
	}

	config, _ := config.LoadConfig(".")

	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPrivateKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
		return
	}

	user, err := handler.userUsecase.FindById(fmt.Sprint(sub))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		return
	}

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "", false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "", false, false)

	c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}

func (handler *AuthHandler) LogoutUser(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.SetCookie("logged_in", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
