package http

import (
	"chat/services/domain/user/usecase"
	"github.com/gin-gonic/gin"
)

type MessageRouter struct {
	messageHandler MessageHandler
}

func NewRoomMessageRouter(messageHandler MessageHandler) MessageRouter {
	return MessageRouter{messageHandler}
}

func (m MessageRouter) MessageRoute(rg *gin.RouterGroup, userUseCase usecase.UserUseCase) {
	router := rg.Group("/message")
	router.POST("/:room_id", m.messageHandler.AddMessage)
}
