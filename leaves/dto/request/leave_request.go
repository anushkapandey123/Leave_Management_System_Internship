package request

type LeaveRequest struct {
	Id int `json:"id"`
	StartDate string `json:"startdate"`
	EndDate string `json:"enddate"`
}