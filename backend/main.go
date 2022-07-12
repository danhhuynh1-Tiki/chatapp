package main

import (
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"

	repository "jwt-demo/repository/mongo"
	routes "jwt-demo/routes"
)

const connString = "localhost:8000"

func main() {
	router := gin.Default()
	db := repository.ConnectDatabase() // connect to MongoDB and keep the connection

	if db == nil { // can not connect to the database
		return
	}

	router = routes.UserRoutes(router, db)

	if err := router.Run(connString); err != nil {
		fmt.Println("Can not run the server...")
	}
}
