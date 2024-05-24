package service

import (
	"context"
	"errors"

	"reflect"
	"time"

	"main.go/leaves/dto/request"
	"main.go/leaves/model"
)

type EmployeeRepository interface {
	FindByUserId(context.Context) (*[]model.Emp, error)
	Create(context.Context, *model.Leave) error
	FetchLeavesByEmpId(context.Context) (*[]model.Leave, error)
	// FindByStartDate(context.Context, time.Time) error
}

type employeeService struct {
	empRepo EmployeeRepository
}

func NewEmployeeService(emprepo EmployeeRepository) *employeeService {
	return &employeeService{
		empRepo: emprepo,
	}
}

func (c *employeeService) EmployeeDetails(ctx context.Context) (*[]model.Emp, error) {
	employee, err := c.empRepo.FindByUserId(ctx)
	if err != nil {
		// return nil, ae.InternalServerError("Something went wrong", " ", fmt.Errorf("%v", err))
		return nil, errors.New("error occured")
	}

	emptyEmployee := []model.Emp{}

	if err == nil && reflect.DeepEqual(employee, emptyEmployee) {
		err := errors.New("user does not exist")
		// return nil, ae.BadRequestError("Bad request error", " ", fmt.Errorf("%v", err))
		return nil, err
	}

	return employee, nil
}

func (c *employeeService) InsertLeave(ctx context.Context, newLeaveRequest request.LeaveRequest) error {
	sd := newLeaveRequest.StartDate
	ed := newLeaveRequest.EndDate

	layout := "2006-01-02"
	sdate, _ := time.Parse(layout, sd)
	edate, _ := time.Parse(layout, ed)

	if edate.Before(sdate) {
		return errors.New("bad request")
	}
	leave := model.Leave{EmpId: newLeaveRequest.Id, StartDate: sdate, EndDate: edate}
	err1 := c.empRepo.Create(ctx, &leave)
	if err1 != nil {

		return err1
	}
	return nil
}

func (c *employeeService) LeaveDetailsOfMembers(ctx context.Context) (*[]model.Leave, error) {
	leave, err := c.empRepo.FetchLeavesByEmpId(ctx)

	if err != nil {
		// return nil, ae.InternalServerError("Something went wrong", " ", fmt.Errorf("%v", err))
		return nil, errors.New("error occured")
	}

	emptyLeave := model.Leave{}

	if err == nil && reflect.DeepEqual(leave, emptyLeave) {
		err := errors.New("user does not exist")
		// return nil, ae.BadRequestError("Bad request error", " ", fmt.Errorf("%v", err))
		return nil, err
	}

	return leave, nil



}

// func (c *employeeService) DeleteLeave(ctx context.Context, newDeleteLeaveRequest request.DeleteLeaveRequest) error {
// 	sd := newDeleteLeaveRequest.StartDate
// 	layout := "2006-01-02"
// 	sdate, _ := time.Parse(layout, sd)

// 	err1 := c.empRepo.FindByStartDate(ctx, sdate)

// 	if err1 != nil {
// 		return err1
// 	}

// 	return nil

// }
