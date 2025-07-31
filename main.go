package main

import (
	"lovablytics/cmd/server"
	"lovablytics/cmd/server/config"
)

func main() {
	config.LoadEnv()
	server.Start()
}
