package http

import (
	"github.com/gin-gonic/gin"

	middleware "chat/pkg/middleware"
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
	router.GET("/me", uc.userHandler.GetMe)
	//router.GET("/", uc.userHandler.GetAll)
	router.GET("/", uc.userHandler.FilterUser)
}
