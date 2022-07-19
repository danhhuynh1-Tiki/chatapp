package main

import (
	"chat/database"
	"fmt"

	// "fmt"
	"time"

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
	"github.com/go-co-op/gocron"
)

func main() {
	app := useRoutes()
	go CheckStatus()
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

	return r

}

func CheckStatus() {
	DB := database.GetDatabase()
	s := gocron.NewScheduler(time.UTC)
	userRepo := userRepository.NewRepository(DB)

	s.Every(10).Seconds().Do(func() {
		listUser := userRepo.GetAllUser()
		for i := range listUser {
			if time.Now().UTC().Unix()-listUser[i].Request_At.Unix() > 5 {
				err := userRepo.UpdateStatusUser(listUser[i].ID, listUser[i].Request_At, 0)
				if err != nil {
					fmt.Println("Cannot update status", listUser[i].ID)
				}
			}
		}
	})

	s.StartBlocking()

}
