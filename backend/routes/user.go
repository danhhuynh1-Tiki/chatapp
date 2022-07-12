package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	controllers "jwt-demo/controllers"
	usecase "jwt-demo/usecase"
	repository "jwt-demo/repository/mongo"
)


var UserRoutes = func(router *gin.Engine, db *mongo.Client) *gin.Engine {
	database := repository.NewUserMongo(db)
	service := usecase.NewUserService(database)
	controller := controllers.NewUserController(service)

	router.POST("/users", controller.Create)
	router.POST("/users/login", controller.Login)

	return router
}