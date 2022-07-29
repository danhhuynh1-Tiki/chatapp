package http

import (
	middleware "chat/pkg/middleware"
	//"chat/services/domain/room/usecase"
	usecase "chat/services/domain/user/usecase"
	"github.com/gin-gonic/gin"
)

type RoomRouter struct {
	roomHandler RoomHandler
}

func NewRoomRouter(roomHandler RoomHandler) RoomRouter {
	return RoomRouter{roomHandler}
}
func (r *RoomRouter) RoomRoute(rg *gin.RouterGroup, userUseCase usecase.UserUseCase) {
	router := rg.Group("/room")
	router.Use(middleware.DeserializeUser(userUseCase))
	router.GET("/:id", r.roomHandler.CreateRoom)
	// Create Group
	router.POST("/", r.roomHandler.CreateGroup)
	router.GET("/group/:email", r.roomHandler.GetGroup)
}
