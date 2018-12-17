package main

import (
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"net/http"
)

// here I provide app token from test chat
// please contact @Pavel Maksymov if u want to join channel or use another token
var verificationToken = "xoxb-463453827013-463161778324-pSVy6SbQBHQmQaKgTheryh01"

func main() {
	// test events functionality
	http.HandleFunc("/slash", func(w http.ResponseWriter, r *http.Request) {
		s, err := slack.SlashCommandParse(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !s.ValidateToken(verificationToken) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch s.Command {
		case "/echo":
			params := &slack.Msg{Text: s.Text}
			b, err := json.Marshal(params)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":18080", nil)
}
