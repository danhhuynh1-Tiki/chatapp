package http

import (
	"chat/services/domain/message/usecase"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	messageUseCase usecase.MessageUseCase
}

func NewMessageHandler(messageUseCase usecase.MessageUseCase) MessageHandler {
	return MessageHandler{messageUseCase: messageUseCase}
}

func (m *MessageHandler) AddMessage(c *gin.Context) {

}
func (m *MessageHandler) GetMessage(c *gin.Context) {

}
