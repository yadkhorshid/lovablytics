package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Track struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Artist string   `json:"artist"`
	Genres []string `json:"genres"`
	Mood   string   `json:"mood"`
}

type AudioFeature struct {
	ID           string  `json:"id"`
	Danceability float64 `json:"danceability"`
	Energy       float64 `json:"energy"`
	Valence      float64 `json:"valence"`
}

type EnrichedTrack struct {
	Track
	AudioFeature
}

func ExchangeSpotifyCode(code string) ([]byte, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", os.Getenv("SPOTIFY_REDIRECT_URI"))
	data.Set("client_id", os.Getenv("SPOTIFY_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("SPOTIFY_CLIENT_SECRET"))

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, errors.New("failed to exchange token with Spotify: " + string(body))
	}
	return io.ReadAll(resp.Body)
}

// Deprecated API for getting audio features so had to create my own switch case
func InferMood(genres []string) string {
	for _, g := range genres {
		switch {
		case strings.Contains(g, "dance"), strings.Contains(g, "pop"):
			return "upbeat"
		case strings.Contains(g, "chill"), strings.Contains(g, "ambient"):
			return "relaxed"
		case strings.Contains(g, "emo"), strings.Contains(g, "punk"):
			return "emotional"
		case strings.Contains(g, "hip hop"), strings.Contains(g, "trap"), strings.Contains(g, "rap"):
			return "intense"
		case strings.Contains(g, "acoustic"), strings.Contains(g, "folk"):
			return "calm"
		case strings.Contains(g, "rock"), strings.Contains(g, "alternative"), strings.Contains(g, "metal"):
			return "energetic"
		case strings.Contains(g, "jazz"), strings.Contains(g, "blues"):
			return "smooth"
		}
	}
	return "unknown"
}

func FetchTopTracksWithMood(authHeader string) ([]Track, error) {
	client := &http.Client{}

	// Step 1: Get user's top tracks
	req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/me/top/tracks?limit=10", nil)
	req.Header.Add("Authorization", authHeader)

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println("Top Tracks Error:", string(bodyBytes))
		return nil, errors.New("failed to fetch top tracks")
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data struct {
		Items []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Artists []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"artists"`
		} `json:"items"`
	}
	json.Unmarshal(body, &data)

	var tracks []Track

	// Step 2: For each track, fetch artist's genres
	for _, item := range data.Items {
		artistID := item.Artists[0].ID
		artistName := item.Artists[0].Name

		artistReq, _ := http.NewRequest("GET", "https://api.spotify.com/v1/artists/"+artistID, nil)
		artistReq.Header.Add("Authorization", authHeader)
		artistResp, err := client.Do(artistReq)
		if err != nil || artistResp.StatusCode != 200 {
			continue
		}
		defer artistResp.Body.Close()
		artistBody, _ := io.ReadAll(artistResp.Body)

		var artistData struct {
			Genres []string `json:"genres"`
		}
		json.Unmarshal(artistBody, &artistData)

		mood := InferMood(artistData.Genres)

		tracks = append(tracks, Track{
			ID:     item.ID,
			Name:   item.Name,
			Artist: artistName,
			Genres: artistData.Genres,
			Mood:   mood,
		})
	}

	return tracks, nil
}
