package bot

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"strings"
)

var commands = map[string]func(update Update){
	"/start": Start,
	"/help":  Help,
}

type Message struct {
	MessageID int `json:"message_id"`
	From      struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	} `json:"from"`
	Chat struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	} `json:"chat"`
	Date int    `json:"date"`
	Text string `json:"text"`
}
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

func BotStart() {
	url := viper.GetString("token")

	for {
		resp, err := http.Get(url + "getUpdates")
		if err != nil {
			fmt.Println("Error getting updates:", err)
			continue
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}
		resp.Body.Close()
		var updates []Update
		err = json.Unmarshal(body, &updates)
		if err != nil {
			fmt.Println("Error unmarshalling updates:", err)
			continue
		}
		for _, update := range updates {
			if update.Message != (Message{}) && len(update.Message.Text) > 0 && update.Message.Text[0] == '/' {
				command := strings.SplitN(update.Message.Text, " ", 2)[0]
				if handler, ok := commands[command]; ok {
					handler(update)
				} else {
					SendMessage(update.Message.Chat.ID, fmt.Sprintf("Unknown command: %s! type '/help' to get the list of commands", command))
				}
			}
		}

	}
}
