package http

import (
	"chat/pkg/config"
	"chat/pkg/utils"
	"chat/services/domain/room/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomHandler struct {
	roomUsecase usecase.RoomUsecase
}

func NewRoomHandler(roomUsecase usecase.RoomUsecase) RoomHandler {
	return RoomHandler{roomUsecase}
}

func (u *RoomHandler) CreateRoom(c *gin.Context) {
	user_id1 := c.Param("id")
	cookie, _ := c.Cookie("access_token")

	config, _ := config.LoadConfig(".")

	id, _ := utils.ValidateToken(cookie, config.AccessTokenPublicKey)

	user_id, _ := primitive.ObjectIDFromHex(fmt.Sprint(id))

	fmt.Println(user_id1)
	fmt.Println(user_id)

	fmt.Println(id)

}
func (u *RoomHandler) CreateGroup(c *gin.Context) {

}
func (u *RoomHandler) GetRoom(c *gin.Context) {

}
