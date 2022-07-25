package http

import (
	"chat/pkg/config"
	"chat/services/domain/user/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"

	filter "chat/pkg/filter"
	"chat/pkg/utils"
	"chat/services/models"
)

type UserHandler struct {
	usecase usecase.UserUseCase
}

func NewUserHandler(handler usecase.UserUseCase) UserHandler {
	return UserHandler{
		usecase: handler,
	}
}

func (handler *UserHandler) GetMe(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(*models.DBResponse)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   gin.H{"user": filter.FilteredResponse(currentUser)},
	})
}
func (handler *UserHandler) GetAll(c *gin.Context) {
	listUser, err := handler.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot get user",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    listUser,
	})
}
func (handler *UserHandler) FilterUser(c *gin.Context) {
	str := c.Query("size")
	cookie, _ := c.Cookie("access_token")
	config, _ := config.LoadConfig(".")

	id, _ := utils.ValidateToken(cookie, config.AccessTokenPublicKey)
	user_id, _ := primitive.ObjectIDFromHex(fmt.Sprint(id))
	i, _ := strconv.ParseInt(str, 10, 64)
	//
	//fmt.Println("count  : ", i)
	//
	filter, err := handler.usecase.FilterUser(user_id, i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    filter,
		"message": "successfull",
	})
}
