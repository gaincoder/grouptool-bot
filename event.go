package main

import (
	"os"
	"time"
)

type Event struct {
	Id       string
	Name     string
	Location string
	Date     time.Time
	Public   bool
}

func (e Event) Url() string {
	url := os.Getenv("PORTAL_URL") + "/event/view/" + e.Id
	return url
}
