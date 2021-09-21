package main

import (
	"os"
)

type Poll struct {
	Id   string
	Name string
}

func (p Poll) Url() string {
	url := os.Getenv("PORTAL_URL") + "/poll/view/" + p.Id
	return url
}
