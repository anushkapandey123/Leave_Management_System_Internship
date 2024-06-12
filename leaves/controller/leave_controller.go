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

type LeaveService interface {
	InsertLeave(context.Context, request.LeaveRequest) error
	LeaveDetailsOfMembers(context.Context) (*[]model.Leave, error)
	DeleteLeave(context.Context, request.DeleteLeaveRequest) error
}

type LeaveController struct {
	leaveService LeaveService
}

func NewLeaveController(leaveservice LeaveService) *LeaveController {
	return &LeaveController{
		leaveService: leaveservice,
	}
}

func (ec *LeaveController) Insert(c *gin.Context) {
	var newLeaveRequest request.LeaveRequest

	fmt.Println("inside controller")

	if err := c.BindJSON(&newLeaveRequest); err != nil {
		errors.New("error occured, bad request")
		c.AbortWithStatus(404)
	}

	fmt.Println(newLeaveRequest)

	responseError := ec.leaveService.InsertLeave(c.Request.Context(), newLeaveRequest)

	if responseError != nil {

		errors.New("insertion failed")
		c.AbortWithStatus(404)

	}
	c.JSON(http.StatusOK, "Leave Insertion Successful")

}

func (ec *LeaveController) LeaveDetails(c *gin.Context) {

	leave, err := ec.leaveService.LeaveDetailsOfMembers(c.Request.Context())

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, *leave)

}

func (ec *LeaveController) Delete(c *gin.Context) {
	var newDeleteLeaveRequest request.DeleteLeaveRequest

	if err := c.BindJSON(&newDeleteLeaveRequest); err != nil {
		errors.New("error occured, bad request")
		return
	}

	responseError := ec.leaveService.DeleteLeave(c.Request.Context(), newDeleteLeaveRequest)

	if responseError != nil {

		errors.New("deletion failed")
		c.AbortWithStatus(404)
	}
	c.JSON(http.StatusOK, "Leave Deletion Successful")

}
