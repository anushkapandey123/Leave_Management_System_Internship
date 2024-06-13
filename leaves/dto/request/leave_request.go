package request

type LeaveRequest struct {
	// Id int `json:"empId"`
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
	LeaveType string `json:"leaveType"`
}