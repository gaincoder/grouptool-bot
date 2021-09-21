package main

import (
	"bytes"
	"text/template"

	"github.com/enescakir/emoji"
)

type BirthdayInTwoWeeksEvent struct {
	Birthday Birthday
}

func (b BirthdayInTwoWeeksEvent) Message(bot *Bot) {
	tmpl := template.Must(template.ParseFiles("templates/birthday_two_weeks.tpl"))
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, b)
	text := tpl.String()
	emojiText := emoji.Parse(text)
	bot.ApiInput(emojiText)
}
