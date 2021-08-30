package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"main.go/Week_5/Homework/RestAPI/config"
	"main.go/Week_5/Homework/RestAPI/controllers"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/movie/:slug", inDB.GetMovie)
	router.GET("/movie", inDB.GetMovie)
	router.POST("/movie", inDB.CreateMovie)
	router.PUT("/movie/:slug", inDB.UpdateMovie)
	router.DELETE("/movie/:slug", inDB.DeleteMovie)
	router.Run(":8080")
}
