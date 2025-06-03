package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type DiscordClient struct {
	Session *discordgo.Session
}

func NewDiscordClient(token string) (*DiscordClient, error) {
	log.Print("Initializing Discord client with token: ", token)
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	err = session.Open()
	if err != nil {
		return nil, err
	}

	log.Println("Discord session established")
	return &DiscordClient{Session: session}, nil
}

func (dc *DiscordClient) SendMessage(channelID, message string) error {
	_, err := dc.Session.ChannelMessageSend(channelID, message)
	return err
}

func (dc *DiscordClient) AddRole(memberID, roleID, guildID string) error {
	err := dc.Session.GuildMemberRoleAdd(guildID, memberID, roleID)
	return err
}

func (dc *DiscordClient) Close() {
	dc.Session.Close()
}
