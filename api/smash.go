package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/slack-go/slack"
)

func HandleSmashCommand(w http.ResponseWriter, r *http.Request) {
	// todo: verify signing secret

	defer r.Body.Close()
	cmd, err := slack.SlashCommandParse(r)
	if err != nil {
		log.Printf("[error] parse slash command: %+v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("User '%s' is creating an arena", cmd.UserName)
	cta := fmt.Sprintf("@here: *%s* is creating a new online Smash arena. Join in!", cmd.UserName)
	if cmd.Text != "" {
		cta = fmt.Sprintf("@here: *%s* is creating a new Smash arena at *%s*. Join in!", cmd.UserName, cmd.Text)
	}

	headerText := slack.NewTextBlockObject("mrkdwn", cta, false, false)
	joinText := slack.NewTextBlockObject("plain_text", "Join!", false, false)
	joinBtn := slack.NewButtonBlockElement("", "join", joinText)

	msg := slack.NewBlockMessage(
		slack.NewSectionBlock(headerText, nil, nil),
		slack.NewActionBlock("", joinBtn),
	)
	msg.ResponseType = "in_channel"

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		log.Printf("[error] respond to command: %+v", err)
		return
	}
}
