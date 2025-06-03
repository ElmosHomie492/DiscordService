package config

import (
	"errors"
	"log"
	"os"
)

type Config struct {
	DiscordToken string
	Port         string
}

func LoadConfig() (*Config, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("missing Discord token")
	}

	discordToken := os.Args[1]
	log.Print(discordToken)

	return &Config{
		DiscordToken: discordToken,
		Port:         "8080",
	}, nil
}
