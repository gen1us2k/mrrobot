package config

import "github.com/kelseyhightower/envconfig"

const (
	// EnvProduction is a production environment
	EnvProduction = "production"
	// EnvDevelopment is a development environment
	EnvDevelopment = "development"
)

// BotConfig is a struct that stores configuration parsed by `envconfig`
// environment variables
type BotConfig struct {
	Env            string `envconfig:"ENV" default:"development"`
	BindAddr       string `envconfig:"BIND_ADDR" default:":12022"`
	SigningSecret  string `envconfig:"SLACK_SIGNING_SECRET"`
	SlackBotToken  string `envconfig:"SLACK_BOT_TOKEN"`
	WelcomeMessage string `envconfig:"WELCOME_MESSAGE"`
}

// Parse parses and returns BotConfig structure
func Parse() (*BotConfig, error) {
	var c BotConfig
	err := envconfig.Process("", &c)
	return &c, err
}
