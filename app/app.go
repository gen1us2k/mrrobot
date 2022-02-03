package app

import (
	"encoding/json"
	"greeter_bot/config"
	"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
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
	http.HandleFunc("/", h.handle)
}

func (l *LambdaHandler) Init(c *config.BotConfig) {
	l.config = c
}
func (h *HTTPHandler) handle(w http.ResponseWriter, r *http.Request) {
	var api = slack.New(h.config.SlackBotToken)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sv, err := slack.NewSecretsVerifier(r.Header, h.config.SigningSecret)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, err := sv.Write(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := sv.Ensure(); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}
	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		spew.Dump(innerEvent)
		switch ev := innerEvent.Data.(type) {
		case *slackevents.TeamJoinEvent:
			_ = ev
			spew.Dump(api.PostMessage(ev.User.ID, slack.MsgOptionText("Yes, hello.", false)))
		}
	}
}
func (h *HTTPHandler) Start() error {
	return http.ListenAndServe(":12022", nil)
}
