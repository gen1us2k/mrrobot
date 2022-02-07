package app

import (
	"encoding/json"
	"greeter_bot/config"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/akrylysov/algnhsa"
	"github.com/davecgh/go-spew/spew"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type (
	// Handler is an interface for the webserver that handles
	// incoming requests from Slack events API
	//
	// You can add support of any cloud provider by implementing this interface
	Handler interface {
		Init(c *config.BotConfig)
		Start() error
	}
	// HTTPHandler is an implementation of webserver for local development/testing
	HTTPHandler struct {
		Handler
		config *config.BotConfig
	}
)

// NewHandler creates slack events api handler
// It creates HTTPHandler for development environment
// and LambdaHandler for production env
func NewHandler(c *config.BotConfig) Handler {
	var h Handler
	h = &HTTPHandler{}
	h.Init(c)
	return h

}

// Init initializes handler
func (h *HTTPHandler) Init(c *config.BotConfig) {
	h.config = c
	http.HandleFunc("/", h.handle)
}

// handle handles incoming data from
func (h *HTTPHandler) handle(w http.ResponseWriter, r *http.Request) {
	var api = slack.New(h.config.SlackBotToken)
	body, err := ioutil.ReadAll(r.Body)
	spew.Dump("parse body")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	spew.Dump("read body")
	spew.Dump(r.Header)
	sv, err := slack.NewSecretsVerifier(r.Header, h.config.SigningSecret)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	spew.Dump(string(body))
	spew.Dump("read secrets verifier")
	if _, err := sv.Write(body); err != nil {
		spew.Dump(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := sv.Ensure(); err != nil {
		spew.Dump(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	spew.Dump("ensure")
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
			text := strings.Replace(h.config.WelcomeMessage, "\\n", "\n", -1)
			spew.Dump(api.PostMessage(ev.User.ID, slack.MsgOptionText(text, false)))
		}
	}
}

// Start starts the server
func (h *HTTPHandler) Start() error {
	if h.config.Env == config.EnvDevelopment {
		return http.ListenAndServe(h.config.BindAddr, nil)
	}
	algnhsa.ListenAndServe(http.DefaultServeMux, nil)
	return nil
}
