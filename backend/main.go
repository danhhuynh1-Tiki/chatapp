package main

import (
	"chat/database"
	// user
	userHandler "chat/user/delivery/http"
	userRepository "chat/user/repository"
	userUsecase "chat/user/usecase"

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
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	userHandler.NewUserHandler(r, userUsecase)
	// login
	loginUsecase := loginUsecase.NewLoginUsecase(userRepo)
	loginHandler.NewLoginHandler(r, loginUsecase)
	// signup
	signUsecase := signUsecase.NewSignupUsecase(userRepo)
	singHandler.NewSignupHandler(r, signUsecase)

	return r

}
