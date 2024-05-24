package model

import "time"

type Leave struct {
	EmpId     int       `json:"empid"`
	StartDate time.Time `json:"startdate"`
	EndDate   time.Time `json:"enddate"`
}

func NewLeave(sd time.Time, ed time.Time) Leave {
	return Leave{
		StartDate: sd,
		EndDate:   ed,
	}
}

func (Leave) TableName() string {
	return "leave"
}
