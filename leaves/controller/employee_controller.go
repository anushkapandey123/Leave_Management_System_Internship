package controller

import (
	"context"
	"errors"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/leaves/dto/request"
	"main.go/leaves/model"
)

type EmployeeService interface {
	EmployeeDetails(context.Context) (*[]model.Emp, error)
	InsertLeave(context.Context, request.LeaveRequest) error
	LeaveDetailsOfMembers(context.Context) (*[]model.Leave, error)
	// DeleteLeave(context.Context, request.DeleteLeaveRequest) error
}

type EmployeeController struct {
	employeeService EmployeeService
}

func NewEmployeeController(empservice EmployeeService) *EmployeeController {
	return &EmployeeController{
		employeeService: empservice,
	}
}

// detail
func (ec *EmployeeController) Detail(c *gin.Context) {
	// username, _, _ := c.Request.BasicAuth()
	// user, err := ec.userService.UserDetails(c.Request.Context(), username) // get user details
	// if err != nil {
	// 	appError := err.(*ae.AppError)
	// 	c.AbortWithStatusJSON(appError.HTTPCode(), appError)
	// 	return
	// }

	employee, _ := ec.employeeService.EmployeeDetails(c.Request.Context()) // get customer detail
	// if resError != nil {
	// 	appError := resError.(*ae.AppError)
	// 	c.AbortWithStatusJSON(appError.HTTPCode(), appError)
	// 	return
	// }

	c.JSON(http.StatusOK, *employee)
	// fmt.Println("anushka")
	// c.JSON(http.StatusOK, "hello")
}

func (ec *EmployeeController) Insert(c *gin.Context) {
	var newLeaveRequest request.LeaveRequest

	if err := c.BindJSON(&newLeaveRequest); err != nil {
		errors.New("error occured, bad request")
		return
	}

	responseError := ec.employeeService.InsertLeave(c.Request.Context(), newLeaveRequest)

	if responseError != nil {
		// logger.Error("error occurred. %v", responseError)
		// err := responseError.(*ae.AppError)
		// logger.Error("error occurred. %v", err)
		// c.AbortWithStatusJSON(err.HTTPCode(), err)
		errors.New("insertion failed")
		return

	}
	c.JSON(http.StatusOK, "Leave Insertion Successful")

}

func (ec *EmployeeController) LeaveDetails(c *gin.Context) {

	leave, err := ec.employeeService.LeaveDetailsOfMembers(c.Request.Context())

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, *leave)

}

// func (ec *EmployeeController) Delete(c *gin.Context) {
// 	var newDeleteLeaveRequest request.DeleteLeaveRequest

// 	if err := c.BindJSON(&newDeleteLeaveRequest); err != nil {
// 		errors.New("error occured, bad request")
// 		return
// 	}

// 	responseError := ec.employeeService.DeleteLeave(c.Request.Context(), newDeleteLeaveRequest)

// 	if responseError != nil {
// 		// logger.Error("error occurred. %v", responseError)
// 		// err := responseError.(*ae.AppError)
// 		// logger.Error("error occurred. %v", err)
// 		// c.AbortWithStatusJSON(err.HTTPCode(), err)
// 		errors.New("deletion failed")
// 		return
// 	}
// 	c.JSON(http.StatusOK, "Leave Deletion Successful")

// }
