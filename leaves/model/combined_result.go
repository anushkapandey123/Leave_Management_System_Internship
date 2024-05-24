package model

import "time"

type Result struct {
	Id int
	Name string
	StartDate time.Time
	EndDate time.Time
}