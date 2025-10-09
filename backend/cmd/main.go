package main

import (
	"context"
	"log"

	"main/cmd/api"
	"main/config"

	"github.com/joho/godotenv"
)

// @title Nexus API
// @version 1.0
// @description API for Nexus Project
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error: .env file not found")
	}

	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Fatal error connecting to database: %v", err)
	}
	defer config.DB.Disconnect(context.TODO())

	err = api.RunAPIServer()
	if err != nil {
		log.Fatal(err)
	}
}
