package http

import (
	"chat/domain"
	"chat/jwt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type StatusHandler struct {
	statusUsecase domain.StatusUsecase
}

func NewStatusHandler(router *gin.RouterGroup, statusUsecase domain.StatusUsecase) {

	s := StatusHandler{statusUsecase: statusUsecase}

	// router.GET("/status", s.UpdateStatus)
	router.GET("/message", s.GetMessage)
}

func (s *StatusHandler) GetMessage(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "cannot get cookie",
		})
	}
	token := cookie.Value
	claims, err := jwt.Decode(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "cannot get Token",
		})
	}
	t := time.Now()
	s.statusUsecase.UpdateStatusUser(claims.ID, t, 1)
}
