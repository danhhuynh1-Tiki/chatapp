package http

import (
	"chat/pkg/config"
	"chat/pkg/utils"
	"chat/services/domain/message/usecase"
	"chat/services/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type MessageHandler struct {
	messageUseCase usecase.MessageUseCase
}

func NewMessageHandler(messageUseCase usecase.MessageUseCase) MessageHandler {
	return MessageHandler{messageUseCase: messageUseCase}
}

func (m *MessageHandler) GetMessage(c *gin.Context) {
	room_id, _ := primitive.ObjectIDFromHex(c.Param("room_id"))

	roomM, err := m.messageUseCase.GetMessage(room_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot get message",
		})
		return
	}
	c.JSON(http.StatusOK, roomM)

}

func (m *MessageHandler) AddMessage(c *gin.Context) {
	room_id := c.Param("room_id")
	Room_id, _ := primitive.ObjectIDFromHex(room_id)
	fmt.Println(Room_id)
	cookie, _ := c.Cookie("access_token")
	config, _ := config.LoadConfig(".")
	//
	id, _ := utils.ValidateToken(cookie, config.AccessTokenPublicKey)
	user_id, _ := primitive.ObjectIDFromHex(fmt.Sprint(id))
	fmt.Println("user_id", user_id)
	//
	var message models.Message
	//
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invaild data",
		})
		return
	}
	message.UserID = user_id

	fmt.Println(message)

	roomMessage, err := m.messageUseCase.AddMessage(Room_id, message)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messgae": "Cannot not add message",
		})
		return
	}
	c.JSON(http.StatusOK, roomMessage)
}
