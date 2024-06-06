package controller

import (
	"context"
	"errors"
	"fmt"

	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/leaves/dto/request"
	"main.go/leaves/model"
)

type EmployeeService interface {
	// EmployeeDetails(context.Context) (*[]model.Emp, error)
	InsertLeave(context.Context, request.LeaveRequest) error
	LeaveDetailsOfMembers(context.Context) (*[]model.Leave, error)
	DeleteLeave(context.Context, request.DeleteLeaveRequest) error
}

type EmployeeController struct {
	employeeService EmployeeService
}

func NewEmployeeController(empservice EmployeeService) *EmployeeController {
	return &EmployeeController{
		employeeService: empservice,
	}
}



func (ec *EmployeeController) Insert(c *gin.Context) {
	var newLeaveRequest request.LeaveRequest

	fmt.Println("inside controller")

	if err := c.BindJSON(&newLeaveRequest); err != nil {
		errors.New("error occured, bad request")
		c.AbortWithStatus(404)
	}

	fmt.Println(newLeaveRequest)

	responseError := ec.employeeService.InsertLeave(c.Request.Context(), newLeaveRequest)

	if responseError != nil {
		// logger.Error("error occurred. %v", responseError)
		// err := responseError.(*ae.AppError)
		// logger.Error("error occurred. %v", err)
		// c.AbortWithStatusJSON(err.HTTPCode(), err)
		errors.New("insertion failed")
		c.AbortWithStatus(404)

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

func (ec *EmployeeController) Delete(c *gin.Context) {
	var newDeleteLeaveRequest request.DeleteLeaveRequest

	if err := c.BindJSON(&newDeleteLeaveRequest); err != nil {
		errors.New("error occured, bad request")
		return
	}

	responseError := ec.employeeService.DeleteLeave(c.Request.Context(), newDeleteLeaveRequest)

	if responseError != nil {
		// logger.Error("error occurred. %v", responseError)
		// err := responseError.(*ae.AppError)
		// logger.Error("error occurred. %v", err)
		// c.AbortWithStatusJSON(err.HTTPCode(), err)
		errors.New("deletion failed")
		c.AbortWithStatus(404)
	}
	c.JSON(http.StatusOK, "Leave Deletion Successful")

}
