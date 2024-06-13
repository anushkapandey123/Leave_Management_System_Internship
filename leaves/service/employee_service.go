package service

import (
	"context"
	"errors"
	"fmt"

	// "reflect"
	"time"

	"main.go/leaves/dto/request"
	"main.go/leaves/model"
)

type EmployeeRepository interface {
	FindNameByUserId(context.Context, int) (model.Emp, error)
	Create(context.Context, *model.Leave) error
	FetchLeavesByEmpId(context.Context) (*[]model.Leave, error)
	FindLeave(context.Context, int, time.Time, time.Time) (bool, error)
	Delete(context.Context, *model.Leave, time.Time, time.Time) (error)
	// FindLeaveInARange(context.Context, int, time.Time, time.Time) (bool, error)

	
}

type employeeService struct {
	empRepo EmployeeRepository
}

func NewEmployeeService(emprepo EmployeeRepository) *employeeService {
	return &employeeService{
		empRepo: emprepo,
	}
}



// func (c *employeeService) InsertLeave(ctx context.Context, newLeaveRequest request.LeaveRequest) error {
// 	sd := newLeaveRequest.StartDate
// 	ed := newLeaveRequest.EndDate

// 	layout := "2006-01-02"
// 	sdate, _ := time.Parse(layout, sd)
// 	edate, _ := time.Parse(layout, ed)

// 	if !ValidateLeaveRequest(sdate, edate) {
// 		return errors.New("bad request, end date of the leave cannot be less than start date of the leave")
// 	}

// 	res , err := c.empRepo.FindLeave(ctx, newLeaveRequest.Id, sdate, edate)

// 	fmt.Println(res)

// 	if err != nil {
// 		return err
// 	}

// 	if res == true {

// 		return errors.New("leave record already exists")
// 	}

// 	emp, _ := c.empRepo.FindNameByUserId(ctx, newLeaveRequest.Id);
// 	name := emp.Name;

	

// 	leave := model.Leave{EmpId: uint(newLeaveRequest.Id), Name: name, StartDate: sdate, EndDate: edate, LeaveType: newLeaveRequest.LeaveType, PaidLeavesRemaining: 0, CasualLeavesRemaining: 0}
// 	err1 := c.empRepo.Create(ctx, &leave)
// 	if err1 != nil {

// 		return err1
// 	}
// 	return nil
// }


// func (c *employeeService) LeaveDetailsOfMembers(ctx context.Context) (*[]model.Leave, error) {
// 	leave, err := c.empRepo.FetchLeavesByEmpId(ctx)

// 	if err != nil {
// 		// return nil, ae.InternalServerError("Something went wrong", " ", fmt.Errorf("%v", err))
// 		return nil, errors.New("error occured")
// 	}

// 	emptyLeave := model.Leave{}

// 	if err == nil && reflect.DeepEqual(leave, emptyLeave) {
// 		err := errors.New("user does not exist")
// 		// return nil, ae.BadRequestError("Bad request error", " ", fmt.Errorf("%v", err))
// 		return nil, err
// 	}

// 	return leave, nil



// }

func (c *employeeService) DeleteLeave(ctx context.Context, newDeleteLeaveRequest request.DeleteLeaveRequest) error {
	sd := newDeleteLeaveRequest.StartDate
	ed := newDeleteLeaveRequest.EndDate
	empId := newDeleteLeaveRequest.Id
	layout := "2006-01-02"
	sdate, _ := time.Parse(layout, sd)
	edate, _ := time.Parse(layout, ed)

	if !ValidateLeaveRequest(sdate, edate) {
		return errors.New("bad request, end date of the leave cannot be less than start date of the leave")
	}



	

	res , err := c.empRepo.FindLeave(ctx, empId, sdate, edate)
	fmt.Println(res)

	leave := model.Leave{}

	if err != nil {
		return err
	}

	if res == true {

		err := c.empRepo.Delete(ctx, &leave, sdate, edate)

		if err != nil {
			return err
		}

	}

	
	
	return err

}

// func ValidateLeaveRequest(start_date time.Time, end_date time.Time) (bool) {

// 	return !end_date.Before(start_date) && !start_date.Before(time.Now()) && !end_date.Before(time.Now())
//  }
