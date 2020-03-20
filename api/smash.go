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
	m := &slack.Msg{Text: fmt.Sprintf("%s is creating a Smash Ultimate arena!", cmd.UserName)}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("[error] respond to command: %+v", err)
		return
	}
}
