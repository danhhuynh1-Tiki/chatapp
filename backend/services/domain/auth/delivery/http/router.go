package http

import (
	"github.com/gin-gonic/gin"

	usecase "chat/services/domain/user/usecase"
	middleware "chat/pkg/middleware"
)

type AuthRouter struct {
	authHandler AuthHandler
}

func NewAuthRouter(authHandler AuthHandler) AuthRouter {
	return AuthRouter{authHandler}
}

func (rc *AuthRouter) AuthRoute(rg *gin.RouterGroup, userUseCase usecase.UserUseCase) {
	router := rg.Group("/auth")

	router.POST("/register", rc.authHandler.SignUpUser)
	router.POST("/login", rc.authHandler.SignInUser)
	router.GET("/refresh", rc.authHandler.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(userUseCase), rc.authHandler.LogoutUser)
}