package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

type Service struct {
	DiscordClient *discordgo.Session
}

func NewService(discordClient *discordgo.Session) *Service {
	return &Service{
		DiscordClient: discordClient,
	}
}

func (s *Service) RegisterRoutes(router *gin.Engine) {
	router.POST("/fillout", s.HandleFillout)
}

func (s *Service) HandleFillout(c *gin.Context) {
	var filloutData struct {
		GuildID string `json:"guild_id"`
		UserID  string `json:"user_id"`
		Role    string `json:"role"`
	}

	if err := c.ShouldBindJSON(&filloutData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := s.AddRoleToUser(filloutData.GuildID, filloutData.UserID, filloutData.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role added successfully"})
}

func (s *Service) AddRoleToUser(guildID, userID, role string) error {
	if s.DiscordClient == nil {
		log.Println("Discord client is not initialized")
		return errors.New("discord client is not initialized")
	}

	if guildID == "" || userID == "" || role == "" {
		log.Println("UserID or Role is empty")
		return errors.New("guildID, userID, and role cannot be empty")
	}

	err := s.DiscordClient.GuildMemberRoleAdd(guildID, userID, role)
	if err != nil {
		log.Printf("Failed to add role %s to user %s: %v", role, userID, err)
		return err
	}

	log.Printf("Successfully added role %s to user %s in guild %s", role, userID, guildID)
	return nil
}
