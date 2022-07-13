package main

import (
	"chat/database"
	userRepository "chat/user/repository"

	// Login
	loginHandler "chat/login/delivery/http"
	loginUsecase "chat/login/usecase"

	// Signup

	singHandler "chat/signup/delivery/http"
	signUsecase "chat/signup/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	app := useRoutes()
	app.Run()
}

func useRoutes() *gin.Engine {
	r := gin.Default()
	DB := database.GetDatabase()
	// user
	userRepo := userRepository.NewRepository(DB)

	// login
	loginUsecase := loginUsecase.NewLoginUsecase(userRepo)
	loginHandler.NewLoginHandler(r, loginUsecase)
	// signup
	signUsecase := signUsecase.NewSignupUsecase(userRepo)
	singHandler.NewSignupHandler(r, signUsecase)

	return r

}
