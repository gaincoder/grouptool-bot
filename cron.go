package main

import (
	"log"

	"github.com/robfig/cron"
)

type crontasks struct {
	bot *Bot
	db  *Database
}

func (c crontasks) checkBirthdayToday() {
	res, err := c.db.db.Query("SELECT id,name,birthdate FROM `birthday` WHERE  DATE_FORMAT(birthdate,'2000-%m-%d') = DATE_FORMAT(NOW(),'2000-%m-%d') ORDER BY DATE_FORMAT(birthdate,'2000-%m-%d')")
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var birthday Birthday
		err := res.Scan(&birthday.Id, &birthday.Name, &birthday.Date)

		if err != nil {
			log.Fatal(err)
		}

		var birthdayToday BirthdayTodayEvent
		birthdayToday.Birthday = birthday
		birthdayToday.Message(c.bot)

	}
}

func (c crontasks) checkBirthdayInTwoWeeks() {
	res, err := c.db.db.Query("SELECT id,name,birthdate FROM `birthday` WHERE  DATE_FORMAT(birthdate,'2000-%m-%d') = DATE_FORMAT(DATE_ADD(now(),INTERVAL 2 WEEK),'2000-%m-%d') ORDER BY DATE_FORMAT(birthdate,'2000-%m-%d')")
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var birthday Birthday
		err := res.Scan(&birthday.Id, &birthday.Name, &birthday.Date)

		if err != nil {
			log.Fatal(err)
		}

		var birthdayToday BirthdayInTwoWeeksEvent
		birthdayToday.Birthday = birthday
		birthdayToday.Message(c.bot)

	}
}

func cronInit(bot *Bot, db *Database) {
	var crontasks crontasks
	crontasks.bot = bot
	crontasks.db = db
	c := cron.New()
	c.AddFunc("0 0 7 * *", crontasks.checkBirthdayToday)
	c.AddFunc("0 0 15 * *", crontasks.checkBirthdayInTwoWeeks)
	c.AddFunc("0 0 14 * * SUN", bot.upcommingEventsCronjobHandler)
	go c.Start()

}
