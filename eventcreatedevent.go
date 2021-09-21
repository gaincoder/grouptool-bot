package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/enescakir/emoji"
)

type EventCreatedEvent struct {
	Event Event
	User  User
}

func (e EventCreatedEvent) Message(bot *Bot) {
	tmpl := template.Must(template.ParseFiles("templates/new_event.tpl"))
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, e)
	bot.ApiInput(emoji.Parse(tpl.String()))
}

func handleEventCreated(bot *Bot) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		var created EventCreatedEvent
		json.Unmarshal(bodyBytes, &created)
		created.Message(bot)
	})
}
