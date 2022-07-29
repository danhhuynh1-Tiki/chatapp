package http

import (
	"chat/pkg/config"
	"chat/pkg/utils"
	"chat/services/domain/room/usecase"
	"chat/services/models"
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
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email",
		})
		return
	}
	roomid, err := u.roomUsecase.CreateGroup(group.Name, group.Admin, group.Members)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"room_id": roomid,
	})
	fmt.Println("group member", group)

}
func (u *RoomHandler) GetGroup(c *gin.Context) {
	email := c.Param("email")
	res, err := u.roomUsecase.GetGroup(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesasge": "Cannot get group",
		})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (u *RoomHandler) GetMembers(c *gin.Context) {
	room_id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot get members",
		})
		return
	}
	listEmail, err := u.roomUsecase.GetMembers(room_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot get members",
		})
		return
	}
	c.JSON(http.StatusOK, listEmail)
}

func (u *RoomHandler) RemoveMember(c *gin.Context) {
	room_id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	email := c.Param("email")
	err := u.roomUsecase.RemoveMember(room_id, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot delete memeber",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})

}

func (u *RoomHandler) AddMember(c *gin.Context) {
	room_id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	email := c.Param("email")
	err := u.roomUsecase.AddMember(room_id, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot add Members",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
