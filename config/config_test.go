package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	os.Setenv("SLACK_APP_ID", "something")
	os.Setenv("SLACK_CLIENT_ID", "amazing")
	os.Setenv("SLACK_CLIENT_SECRET", "i guess")
	c, err := Parse()
	assert.NoError(t, err)
	assert.Equal(t, "something", c.SlackAppID)
	assert.Equal(t, "amazing", c.SlackClientID)
	assert.Equal(t, "i guess", c.SlackClientSecret)
}
