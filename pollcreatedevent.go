package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/enescakir/emoji"
)

type PollCreatedEvent struct {
	Poll Poll
	User User
}

func (e PollCreatedEvent) Message(bot *Bot) {
	tmpl := template.Must(template.ParseFiles("templates/new_poll.tpl"))
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, e)
	text := tpl.String()
	emojiText := emoji.Parse(text)
	bot.ApiInput(emojiText)
}

func handlePollCreated(bot *Bot) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		var created PollCreatedEvent
		json.Unmarshal(bodyBytes, &created)
		created.Message(bot)
	})
}
