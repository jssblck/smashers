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

	// Create the interactive response that allows other users to join in
	log.Printf("User '%s' is creating an arena", cmd.UserName)
	headerText := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("**%s** is creating a new Smash arena!", cmd.UserName), false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	joinBtnText := slack.NewTextBlockObject("plan_text", "Join!", false, false)
	joinBtn := slack.NewButtonBlockElement("", "join", joinBtnText)
	actionBlock := slack.NewActionBlock("", joinBtn)

	msg := slack.NewBlockMessage(
		headerSection,
		actionBlock,
	)
	msg.ResponseType = "in_channel"

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		log.Printf("[error] respond to command: %+v", err)
		return
	}
}
