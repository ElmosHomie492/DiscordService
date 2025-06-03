package main

import (
	"fillout_discord_service/internal/api"
	"fillout_discord_service/internal/config"
	"fillout_discord_service/internal/discord"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize router
	router := gin.New()

	// Create a new Discord client
	discordClient, err := discord.NewDiscordClient(cfg.DiscordToken)
	if err != nil {
		log.Fatalf("Error creating Discord client: %v", err)
	}
	log.Println("Discord client initialized successfully")

	// Register the Discord client with the API service
	apiService := api.NewService(discordClient.Session)

	// Register the API routes with the router
	apiService.RegisterRoutes(router)

	// Note: If you have a separate handler for Fillout, you can register it here
	// ex: router.HandleFunc("/fillout", fillout.ProcessFilloutData).Methods("POST")

	// Start the server
	log.Printf("Starting server on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
