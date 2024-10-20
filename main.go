package main

import (
	"golang-unit-test/db"
	"golang-unit-test/entity"
	"golang-unit-test/handler"
	"golang-unit-test/repository"
	"golang-unit-test/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.SetupDatabaseConnection()

	// Auto migrate database
	db.DB.AutoMigrate(&entity.Post{})

	r := gin.Default()

	// Post init
	PostRepository := repository.NewPostRepository(db.DB)
	PostService := service.NewPostService(PostRepository)
	PostHandler := handler.NewPostHandler(PostService)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/create-post", PostHandler.PostCreateHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}
