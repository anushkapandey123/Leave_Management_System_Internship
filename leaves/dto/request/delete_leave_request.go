package request

type DeleteLeaveRequest struct {
	Id int `json:"id"`
	StartDate string `json:"startdate"`
	EndDate string `json:"enddate"`
}