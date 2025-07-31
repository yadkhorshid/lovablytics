package services

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
)

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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
