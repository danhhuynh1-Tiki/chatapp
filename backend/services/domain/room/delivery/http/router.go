package http

import "github.com/gin-gonic/gin"

type RoomRouter struct {
	roomHandler RoomHandler
}

func NewRoomRouter(roomHandler RoomHandler) RoomRouter {
	return RoomRouter{roomHandler}
}
func (r *RoomRouter) RoomRoute(rg *gin.RouterGroup) {
	router := rg.Group("/room")
	router.GET("/:id", r.roomHandler.CreateRoom)
	router.POST("/", r.roomHandler.CreateGroup)

}
