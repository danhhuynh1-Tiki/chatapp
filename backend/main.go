package main

import (
	"chat/database"
	// module user
	userHttp "chat/user/delivery/http"
	userRepo "chat/user/repository"
	userUse "chat/user/usecase"

	// module login

	loginHttp "chat/login/delivery/http"
	loginUse "chat/login/usecase"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := UseRoute()
	err := r.Run()
	if err != nil {
		fmt.Println("Server run Failed")
	}
}

func UseRoute() *gin.Engine {
	r := gin.Default()

	db := database.GetDatabase()

	// user
	userRepo := userRepo.NewRepository(db)
	userUse := userUse.NewUserUseCase(userRepo)

	userHttp.NewUserHandler(r, userUse)

	// login
	loginUsecase := loginUse.NewLoginUsecase(userRepo)
	loginHttp.NewLoginHandler(r, loginUsecase)
	return r
}
