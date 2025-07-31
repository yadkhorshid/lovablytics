package server

import (
	"lovablytics/cmd/server/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/analyze", handlers.Analyze)
	router.GET("/auth/spotify/callback", handlers.SpotifyCallback)
	router.GET("/spotify/profile", handlers.GetSpotifyProfile)

	router.Run(":8080")
}
