package config

type Config struct {
	DiscordToken string
	Port         string
}

func LoadConfig() (*Config, error) {
	return &Config{
		DiscordToken: "MTM3OTI0NTExMjU4NDA0ODcxMQ.G3CPO2.Ikr5-kJNhWUDjLb7IX9X_Y_cird0swa8VVpEEA",
		Port:         "8080",
	}, nil
}
