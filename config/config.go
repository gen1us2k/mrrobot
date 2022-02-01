package config

import "github.com/kelseyhightower/envconfig"

type BotConfig struct {
	SlackAppID        string `envconfig:"SLACK_APP_ID"`
	SlackClientID     string `envconfig:"SLACK_CLIENT_ID"`
	SlackClientSecret string `envconfig:"SLACK_CLIENT_SECRET"`
}

func Parse() (*BotConfig, error) {
	var c BotConfig
	err := envconfig.Process("", &c)
	return &c, err
}
