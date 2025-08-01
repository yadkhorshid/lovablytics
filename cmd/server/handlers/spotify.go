package handlers

import (
	"lovablytics/cmd/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SpotifyCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Spotify code"})
		return
	}

	tokenResp, err := services.ExchangeSpotifyCode(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json", tokenResp)
}

func GetSpotifyProfile(c *gin.Context) {
	accessToken := c.GetHeader("Authorization") // expects "Bearer <token>"

	if accessToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing access token"})
		return
	}

	profile, err := services.GetSpotifyProfile(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
