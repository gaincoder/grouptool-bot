package main

import "os"

type Info struct {
	Id        string
	Headline  string
	Text      string
	Important bool
}

func (e Info) Url() string {
	url := os.Getenv("PORTAL_URL") + "/info"
	return url
}
