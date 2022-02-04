package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	os.Setenv("SLACK_SIGNING_SECRET", "something")
	os.Setenv("SLACK_BOT_TOKEN", "amazing")
	os.Setenv("WELCOME_MESSAGE", "i guess")
	c, err := Parse()
	assert.NoError(t, err)
	assert.Equal(t, "something", c.SigningSecret)
	assert.Equal(t, "amazing", c.SlackBotToken)
	assert.Equal(t, "i guess", c.WelcomeMessage)
	assert.Equal(t, EnvDevelopment, c.Env)
	assert.Equal(t, ":12022", c.BindAddr)
}
