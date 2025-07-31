package handlers

import (
	"lovablytics/cmd/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTopTracksWithMood(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Authorization header"})
		return
	}

	tracks, err := services.FetchTopTracksWithMood(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tracks)
}
