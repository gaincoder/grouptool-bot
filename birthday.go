package main

import "time"

type Birthday struct {
	Id   int32
	Name string
	Date time.Time
}

func (b Birthday) CurrentAge() int {
	now := time.Now()
	years := now.Year() - b.Date.Year()

	birthDay := b.getAdjustedBirthDay()
	if now.YearDay() < birthDay {
		years -= 1
	}

	return years
}

func (b Birthday) NextAge() int {
	return b.CurrentAge() + 1
}

func (b Birthday) getAdjustedBirthDay() int {
	birthDate := b.Date
	now := time.Now()
	birthDay := birthDate.YearDay()
	currentDay := now.YearDay()
	if isLeap(birthDate) && !isLeap(now) && birthDay >= 60 {
		return birthDay - 1
	}
	if isLeap(now) && !isLeap(birthDate) && currentDay >= 60 {
		return birthDay + 1
	}
	return birthDay
}

func isLeap(date time.Time) bool {
	year := date.Year()
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}
