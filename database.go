package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func connectToDatabase() *Database {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_CONNECTION"))

	if err != nil {
		panic(err.Error())
	}

	return &Database{db: db}
}

func (database Database) close() {
	database.db.Close()
}

func (database Database) nextFiveEvents() []Event {

	res, err := database.db.Query("SELECT id,name,location,date,public FROM `event` WHERE `date` >= NOW() AND deleted_at IS NULL ORDER BY date LIMIT 5")
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	events := []Event{}

	for res.Next() {

		var event Event
		err := res.Scan(&event.Id, &event.Name, &event.Location, &event.Date, &event.Public)

		if err != nil {
			log.Fatal(err)
		}

		events = append(events, event)

	}
	return events
}

func (database Database) nextFiveInternalEvents() []Event {

	res, err := database.db.Query("SELECT id,name,location,date,public FROM `event` WHERE `date` >= NOW() AND deleted_at IS NULL AND public = 0 ORDER BY date LIMIT 5")
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	events := []Event{}

	for res.Next() {

		var event Event
		err := res.Scan(&event.Id, &event.Name, &event.Location, &event.Date, &event.Public)

		if err != nil {
			log.Fatal(err)
		}

		events = append(events, event)

	}
	return events
}

func (database Database) allBirthdays() []Birthday {

	res, err := database.db.Query("SELECT id,name,birthdate FROM `birthday` ORDER BY DATE_FORMAT(birthdate,'2000-%m-%d')")
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	allBirthdays := []Birthday{}

	for res.Next() {

		var birthday Birthday
		err := res.Scan(&birthday.Id, &birthday.Name, &birthday.Date)

		if err != nil {
			log.Fatal(err)
		}

		allBirthdays = append(allBirthdays, birthday)

	}
	return allBirthdays
}
