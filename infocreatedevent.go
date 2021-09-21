package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/enescakir/emoji"
)

type InfoCreatedEvent struct {
	Info Info
	User User
}

func (e InfoCreatedEvent) Message(bot *Bot) {
	tmpl := template.Must(template.ParseFiles("templates/new_info.tpl"))
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, e)
	bot.ApiInput(emoji.Parse(tpl.String()))
}

func handleInfoCreated(bot *Bot) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		var created InfoCreatedEvent
		json.Unmarshal(bodyBytes, &created)
		created.Message(bot)
	})
}
