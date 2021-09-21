package main

import (
	"bytes"
	"text/template"

	"github.com/enescakir/emoji"
)

type BirthdayTodayEvent struct {
	Birthday Birthday
}

func (b BirthdayTodayEvent) Message(bot *Bot) {
	tmpl := template.Must(template.ParseFiles("templates/birthday_today.tpl"))
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, b)
	text := tpl.String()
	emojiText := emoji.Parse(text)
	bot.ApiInput(emojiText)
}
