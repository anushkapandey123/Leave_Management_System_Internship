package model

import "time"

type Leave struct {
	EmpId     uint       `json:"empid"`
	Name 	  string 	`json:"name"`
	Email 	  string	`json:"email"`
	StartDate time.Time `json:"startdate" gorm:"type:date"`
	EndDate   time.Time `json:"enddate" gorm:"type:date"`
	LeaveType string	`json:"leavetype"`
	PaidLeavesRemaining int	`json:"remaining_paid_leaves"`
	CasualLeavesRemaining int `json:"remaining_casual_leaves"`
}

func NewLeave(sd time.Time, ed time.Time, lt string) Leave {
	return Leave{
		StartDate: sd,
		EndDate:   ed,
		LeaveType: lt,
	}
}

func (Leave) TableName() string {
	return "leave"
}
