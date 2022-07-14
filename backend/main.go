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

	// Talk
	talkHandler "chat/talk/delivery/http"
	talkRepo "chat/talk/repository"
	talkUsecase "chat/talk/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	app := useRoutes()
	app.Run()
}

func useRoutes() *gin.Engine {
	r := gin.Default()
	api := r.Group("/mychat/v1/")
	DB := database.GetDatabase()
	// user
	userRepo := userRepository.NewRepository(DB)
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	userHandler.NewUserHandler(api, userUsecase)
	// login
	loginUsecase := loginUsecase.NewLoginUsecase(userRepo)
	loginHandler.NewLoginHandler(api, loginUsecase)
	// signup
	signUsecase := signUsecase.NewSignupUsecase(userRepo)
	singHandler.NewSignupHandler(api, signUsecase)
	// talk
	talkRepo := talkRepo.NewTalkRepository(DB)
	talkUse := talkUsecase.NewTalkUsecase(talkRepo)
	talkHandler.NewTalkHandler(api, talkUse)
	return r

}
