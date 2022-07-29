package http

import (
	middleware "chat/pkg/middleware"
	"github.com/gin-gonic/gin"

	usecase "chat/services/domain/user/usecase"
)

type UserRouter struct {
	userHandler UserHandler
}

func NewUserRouter(userHandler UserHandler) UserRouter {
	return UserRouter{userHandler}
}

func (uc *UserRouter) UserRoute(rg *gin.RouterGroup, userUseCase usecase.UserUseCase) {
	router := rg.Group("users")
	router.Use(middleware.DeserializeUser(userUseCase))
	router.GET("/me", uc.userHandler.GetUser)
	//router.GET("/", uc.userHandler.GetAll)
	router.GET("/", uc.userHandler.FilterUser)

}
