package services

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetSpotifyProfile(accessToken string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var profile map[string]interface{}
	if err := json.Unmarshal(body, &profile); err != nil {
		return nil, err
	}

	return profile, nil
}
