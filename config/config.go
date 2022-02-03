package config

import "github.com/kelseyhightower/envconfig"

const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
)

type BotConfig struct {
	Env               string `envconfig:"ENV" default:"development"`
	BindAddr          string `envconfig:"BIND_ADDR" default:":12022"`
	SlackAppID        string `envconfig:"SLACK_APP_ID"`
	SlackClientID     string `envconfig:"SLACK_CLIENT_ID"`
	SlackClientSecret string `envconfig:"SLACK_CLIENT_SECRET"`
	SigningSecret     string `envconfig:"SLACK_SIGNING_SECRET"`
	SlackBotToken     string `envconfig:"SLACK_BOT_TOKEN"`
}

func Parse() (*BotConfig, error) {
	var c BotConfig
	err := envconfig.Process("", &c)
	return &c, err
}
