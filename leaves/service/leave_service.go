package service

import (
	"context"
	"errors"
	"fmt"

	"reflect"
	"time"

	"main.go/leaves/dto/request"
	"main.go/leaves/model"
)

type LeaveRepository interface {
	FindNameByUserId(context.Context, int) (model.Emp, error)
	Create(context.Context, *model.Leave) error
	FetchLeavesByEmpId(context.Context) (*[]model.Leave, error)
	FindLeave(context.Context, int, time.Time, time.Time) (bool, error)
	Delete(context.Context, *model.Leave, time.Time, time.Time) error
	GetLatestLeave(context.Context, int) (model.Leave, error)
	CheckForOverlappingLeaves(context.Context, time.Time, int) (bool, error)
	// FindLeaveInARange(context.Context, int, time.Time, time.Time) (bool, error)

}

type leaveService struct {
	leaveRepo LeaveRepository
}

func NewLeaveService(leaverepo LeaveRepository) *leaveService {
	return &leaveService{
		leaveRepo: leaverepo,
	}
}

func (c *leaveService) InsertLeave(ctx context.Context, newLeaveRequest request.LeaveRequest) error {
	sd := newLeaveRequest.StartDate
	ed := newLeaveRequest.EndDate

	layout := "2006-01-02"
	sdate, _ := time.Parse(layout, sd)
	edate, _ := time.Parse(layout, ed)

	if !ValidateLeaveRequest(sdate, edate) {
		return errors.New("bad request, end date of the leave cannot be less than start date of the leave")
	}

	res, err := c.leaveRepo.FindLeave(ctx, newLeaveRequest.Id, sdate, edate)

	fmt.Println(res)

	if err != nil {
		return err
	}

	if res == true {

		return errors.New("leave record already exists")
	}

	// find duration of leave
	// duration := int(edate.Sub(sdate).Hours() / 24) + 1

	paid_leave := 25
	casual_leave := 25

	// only considering weekdays
	duration := CountWeekDays(edate, sdate)

	latestLeave, err := c.leaveRepo.GetLatestLeave(ctx, newLeaveRequest.Id)

	emptyLeave := model.Leave{}

	if err != nil {

		return errors.New("server error")

	} else if reflect.DeepEqual(emptyLeave, latestLeave) {

		paid_leave, casual_leave, err = CalculateRemainingLeaves(newLeaveRequest, paid_leave, casual_leave, duration)

		if err != nil {
			return errors.New("cannot apply for leaves, duration exceeds max limit")
		}

	} else {

		paid_leave, casual_leave, err = CalculateRemainingLeaves(newLeaveRequest, latestLeave.PaidLeavesRemaining, latestLeave.CasualLeavesRemaining, duration)

		if err != nil {
			return errors.New("cannot apply for leaves, duration exceeds max limit")
		}
	}

	//check for overlapping leaves
	resOverlapStartDate, errOverlapStartDate := c.leaveRepo.CheckForOverlappingLeaves(ctx, sdate, newLeaveRequest.Id)
	fmt.Println("overlaping start date : ", resOverlapStartDate)

	resOverlapEndDate, errOverlapEndDate := c.leaveRepo.CheckForOverlappingLeaves(ctx, edate, newLeaveRequest.Id)
	fmt.Println("overlaping end date : ", resOverlapEndDate)

	if resOverlapStartDate || resOverlapEndDate {

		return errors.New("leaves are overlapping")
	} else if errOverlapStartDate != nil || errOverlapEndDate != nil {
		return errors.New("leaves are overlapping")
	}

	emp, _ := c.leaveRepo.FindNameByUserId(ctx, newLeaveRequest.Id)
	name := emp.Name

	leave := model.Leave{EmpId: newLeaveRequest.Id, Name: name, StartDate: sdate, EndDate: edate, LeaveType: newLeaveRequest.LeaveType, PaidLeavesRemaining: paid_leave, CasualLeavesRemaining: casual_leave}
	err1 := c.leaveRepo.Create(ctx, &leave)
	if err1 != nil {

		return err1
	}
	return nil
}

func (c *leaveService) LeaveDetailsOfMembers(ctx context.Context) (*[]model.Leave, error) {
	leave, err := c.leaveRepo.FetchLeavesByEmpId(ctx)

	if err != nil {

		return nil, errors.New("error occured")
	}

	emptyLeave := model.Leave{}

	if err == nil && reflect.DeepEqual(leave, emptyLeave) {
		err := errors.New("user does not exist")
		return nil, err
	}

	return leave, nil

}

func (c *leaveService) DeleteLeave(ctx context.Context, newDeleteLeaveRequest request.DeleteLeaveRequest) error {
	sd := newDeleteLeaveRequest.StartDate
	ed := newDeleteLeaveRequest.EndDate
	empId := newDeleteLeaveRequest.Id
	layout := "2006-01-02"
	sdate, _ := time.Parse(layout, sd)
	edate, _ := time.Parse(layout, ed)

	if !ValidateLeaveRequest(sdate, edate) {
		return errors.New("bad request, end date of the leave cannot be less than start date of the leave")
	}

	res, err := c.leaveRepo.FindLeave(ctx, empId, sdate, edate)
	fmt.Println(res)

	leave := model.Leave{}

	if err != nil {
		return err
	}

	if res == true {

		err := c.leaveRepo.Delete(ctx, &leave, sdate, edate)

		if err != nil {
			return err
		}

	}

	return err

}

func ValidateLeaveRequest(start_date time.Time, end_date time.Time) bool {

	return !end_date.Before(start_date) && !start_date.Before(time.Now()) && !end_date.Before(time.Now())
}

func CalculateRemainingLeaves(newLeaveRequest request.LeaveRequest, paid_leave int, casual_leave int, duration int) (int, int, error) {

	typeOfLeave := newLeaveRequest.LeaveType
	if typeOfLeave == "paid" {

		if duration > (paid_leave) {
			return 0, 0, errors.New("cannot apply for leave, duration exceeds max limit")
		} else {
			paid_leave = paid_leave - int(duration)
		}

	} else {

		if duration > (casual_leave) {
			return 0, 0, errors.New("cannot apply for leave, duration exceeds max limit")
		} else {
			casual_leave = casual_leave - int(duration)
		}

	}

	return paid_leave, casual_leave, nil
}

func CountWeekDays(edate time.Time, sdate time.Time) int {

	days := 0
	for {
		if edate.Equal(sdate) {
			return days + 1
		}
		if sdate.Weekday() != time.Saturday && sdate.Weekday() != time.Sunday {
			days++
		}
		sdate = sdate.Add(time.Hour * 24)
	}

}
