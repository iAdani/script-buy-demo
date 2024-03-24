package main

import (
	"go-server/server/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	router := gin.Default() // Address: http://localhost:8080

	// User routes
	router.POST("/auth", controllers.AuthUID)
	router.GET("/user/:uid", controllers.GetUser)
	router.POST("/user", controllers.AddUser)
	router.DELETE("/user/:uid", controllers.RemoveUser)

	router.GET("/user/:uid/favorites", controllers.GetFavorites)
	router.POST("/user/:uid/favorites", controllers.AddFavorite)
	router.DELETE("/user/:uid/favorites", controllers.RemoveFavorite)
	router.GET("/trending/:cat", controllers.GetTrending)

	// Main routes
	router.GET("/init", controllers.InitDB)
	router.POST("/search", controllers.Search)

	// Test routes
	router.POST("/demo", controllers.Demo)
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
	})

	err = router.Run()
	if err != nil {
		return
	}
}
