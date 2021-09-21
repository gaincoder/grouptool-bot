package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/enescakir/emoji"
	tb "gopkg.in/tucnak/telebot.v2"
)

/*
	Find Emoji shortcuts here: https://github.com/enescakir/emoji/blob/master/map.go
	Emoj View: http://unicode.org/emoji/charts/full-emoji-list.html
*/

type Bot struct {
	db  *Database
	api *tb.Bot
}

func createBot(db *Database) *Bot {

	b, err := tb.NewBot(tb.Settings{
		Token:     os.Getenv("BOT_TOKEN"),
		Poller:    &tb.LongPoller{Timeout: 10 * time.Second},
		ParseMode: "HTML",
	})

	if err != nil {
		log.Fatal(err)
	}

	bot := &Bot{api: b, db: db}

	b.Handle("/veranstaltungen", bot.upcommingEventsHandler)
	b.Handle("/gruppenveranstaltungen", bot.upcommingInternalEventsHandler)
	b.Handle("/geburtstage", bot.allBirthdaysHandler)
	b.Handle("/msg", bot.messageHandler)
	b.Handle("/portal", bot.portalHandler)
	b.Handle("/hilfe", bot.helpHandler)
	b.Handle("/chatid", bot.getChatId)

	return bot

}

func (bot *Bot) upcommingEventsHandler(m *tb.Message) {

	events := []Event{}

	events = bot.db.nextFiveEvents()

	tmpl := template.Must(template.ParseFiles("templates/upcomming_events.tpl"))
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, events)
	bot.api.Send(m.Chat, emoji.Parse(tpl.String()))
}

func (bot *Bot) upcommingInternalEventsHandler(m *tb.Message) {

	events := []Event{}

	events = bot.db.nextFiveInternalEvents()

	tmpl := template.Must(template.ParseFiles("templates/upcomming_internal_events.tpl"))
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, events)
	bot.api.Send(m.Chat, emoji.Parse(tpl.String()))
}

func (bot *Bot) allBirthdaysHandler(m *tb.Message) {

	allBirthdays := []Birthday{}

	allBirthdays = bot.db.allBirthdays()

	tmpl := template.Must(template.ParseFiles("templates/all_birthdays.tpl"))
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, allBirthdays)
	bot.api.Send(m.Chat, emoji.Parse(tpl.String()))
}

func (bot *Bot) helpHandler(m *tb.Message) {

	tmpl := template.Must(template.ParseFiles("templates/help.tpl"))
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, nil)
	bot.api.Send(m.Chat, tpl.String())
}

func (bot *Bot) portalHandler(m *tb.Message) {
	bot.api.Send(m.Chat, os.Getenv("PORTAL_URL"))
}

func (bot *Bot) getChatId(m *tb.Message) {
	chatId := fmt.Sprintf("%d", m.Chat.ID)
	bot.api.Send(m.Chat, chatId)
}

func (bot *Bot) messageHandler(m *tb.Message) {
	group, err := bot.api.ChatByID(os.Getenv("TARGET_CHAT_ID"))
	if err != nil {
		fmt.Println(err.Error())
	}
	bot.api.Send(group, emoji.Parse(m.Payload))
}

func (bot *Bot) ApiInput(input string) {
	chat, err := bot.api.ChatByID(os.Getenv("TARGET_CHAT_ID"))
	if err != nil {
		fmt.Println(err.Error())
	}
	bot.api.Send(chat, input)
}
