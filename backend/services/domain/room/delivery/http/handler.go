package http

import (
	"chat/pkg/config"
	"chat/pkg/utils"
	"chat/services/domain/room/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type RoomHandler struct {
	roomUsecase usecase.RoomUsecase
}

func NewRoomHandler(roomUsecase usecase.RoomUsecase) RoomHandler {
	return RoomHandler{roomUsecase}
}

func (u *RoomHandler) CreateRoom(c *gin.Context) {
	user_id1, _ := primitive.ObjectIDFromHex(c.Param("id"))
	cookie, _ := c.Cookie("access_token")

	config, _ := config.LoadConfig(".")

	id, _ := utils.ValidateToken(cookie, config.AccessTokenPublicKey)

	user_id, _ := primitive.ObjectIDFromHex(fmt.Sprint(id))

	room_id, err := u.roomUsecase.CreateRoom(user_id, user_id1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"room_id": room_id,
	})
}

// Create Group members

func (u *RoomHandler) CreateGroup(c *gin.Context) {
	var emailUsers []struct{ email string }
	if err := c.ShouldBindJSON(&emailUsers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

}
func (u *RoomHandler) GetRoom(c *gin.Context) {

}
