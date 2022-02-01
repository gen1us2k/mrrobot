package app

import (
	"fmt"
	"greeter_bot/config"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Handler interface {
		Init(c *config.BotConfig)
		Start() error
	}
	HTTPHandler struct {
		Handler
		config *config.BotConfig
		echo   *echo.Echo
	}
	LambdaHandler struct {
		Handler
		config *config.BotConfig
	}
	SlackChallenge struct {
		Type      string `json:"type"`
		Token     string `json:"token"`
		Challenge string `json:"challenge"`
	}
	Response struct {
		Message string `json:"message"`
	}
)

func NewHandler(c *config.BotConfig) Handler {
	var h Handler
	h = &HTTPHandler{}
	if c.Env == config.EnvProduction {
		h = &LambdaHandler{}
	}
	h.Init(c)
	return h

}

func (h *HTTPHandler) Init(c *config.BotConfig) {
	h.config = c
	h.echo = echo.New()
	h.echo.POST("/", h.handle)
}

func (l *LambdaHandler) Init(c *config.BotConfig) {
	l.config = c
}
func (h *HTTPHandler) handle(c echo.Context) error {
	body, _ := ioutil.ReadAll(c.Request().Body)
	fmt.Println(string(body))
	var challenge SlackChallenge
	if err := c.Bind(&challenge); err != nil {
		return c.JSON(http.StatusBadRequest, Response{"failed parsing JSON"})
	}
	return c.JSON(http.StatusOK, challenge)

}
func (h *HTTPHandler) Start() error {
	return h.echo.Start(h.config.BindAddr)
}
