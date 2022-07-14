package http

import (
	"chat/domain"
	"chat/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TalkHandler struct {
	talkUsecase domain.TalkUsecase
}

func NewTalkHandler(r *gin.RouterGroup, talkUsecase domain.TalkUsecase) {
	talkHandl := TalkHandler{talkUsecase}
	// r.GET("/talk", talkHandl.GetTalk)
	r.POST("/talk/:id", talkHandl.AddMessage)
	r.GET("/talk/:id", talkHandl.AddTalk)
}
func (t *TalkHandler) GetTalk(c *gin.Context) {

}

func (t *TalkHandler) AddMessage(c *gin.Context) {
	// cookie, _ := c.Request.Cookie("token")

	// claims, _ := jwt.Decode(cookie.Value)
	var m domain.Message

	talk_id := c.Param("id")

	err := c.ShouldBindJSON(&m)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "failed send message",
		})
	}

	err = t.talkUsecase.AddMessage(talk_id, m)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "failed add message",
		})
	}
}

func (t *TalkHandler) AddTalk(c *gin.Context) {

	var m domain.Message

	err := c.ShouldBindJSON(&m)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "failed data",
		})
		return
	}
	cookie, _ := c.Request.Cookie("token")

	user1 := c.Param("id")

	// fmt.Println(cookie.Value)
	claims, _ := jwt.Decode(cookie.Value)
	fmt.Println(user1, claims.ID)
	// fmt.Println(reflect.TypeOf(user2.ID))

	talk, err := t.talkUsecase.GetTalk(user1, claims.ID)
	if err == nil {
		err = t.talkUsecase.AddTalk(user1, claims.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status_code": http.StatusBadRequest,
				"message":     "failed add talk",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status_code": http.StatusOK,
				"message":     "add talk successful",
			})
		}
	} else {
		// talk := t.talkUsecase.GetTalk(user1, claims.ID)
		// fmt.Println(talk)
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "talk is exists",
			"data":        talk.Messages,
		})
		return
	}

}

// func (t *TalkHandler) GetMessage(c *gin.Context) {

// }
