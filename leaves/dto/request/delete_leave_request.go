package request

type DeleteLeaveRequest struct {
	Id int `json:"empId"`
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
	
}