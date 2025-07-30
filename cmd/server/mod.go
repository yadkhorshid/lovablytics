package main

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.POST("/analyze", func(c *gin.Context) {
        var json struct {
            Text string `json:"text" binding:"required"`
        }

        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Text is required"})
            return
        }

        mood := analyzeMood(json.Text)
        recommendations := recommendSongs(mood)

        c.JSON(http.StatusOK, gin.H{
            "mood":           mood,
            "recommendations": recommendations,
        })
    })

    router.Run(":8080")
}

func analyzeMood(text string) string {
    happyKeywords := []string{"happy", "joy", "excited", "love", "great"}
    sadKeywords := []string{"sad", "tired", "bad", "hate", "angry"}

    if containsAny(text, happyKeywords) {
        return "happy"
    } else if containsAny(text, sadKeywords) {
        return "sad"
    }
    return "neutral"
}

func containsAny(text string, keywords []string) bool {
    for _, kw := range keywords {
        if containsIgnoreCase(text, kw) {
            return true
        }
    }
    return false
}

func containsIgnoreCase(text, substr string) bool {
    textLower := strings.ToLower(text)
    substrLower := strings.ToLower(substr)
    return strings.Contains(textLower, substrLower)
}

func recommendSongs(mood string) []string {
    songs := map[string][]string{
        "happy":  {"Happy - Pharrell Williams", "Can't Stop the Feeling - Justin Timberlake"},
        "sad":    {"Someone Like You - Adele", "Fix You - Coldplay"},
        "neutral": {"Let It Be - The Beatles", "Viva La Vida - Coldplay"},
    }

    if recs, ok := songs[mood]; ok {
        return recs
    }
    return []string{}
}
