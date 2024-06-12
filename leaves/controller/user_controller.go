package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/leaves/dto/request"
	"main.go/leaves/model"
	"main.go/leaves/service"
	security "main.go/middleware/security"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (controller *UserController) SignUp(c *gin.Context) {
	var signUpRequest request.SignupRequest
	if err := c.ShouldBindJSON(&signUpRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Name:        signUpRequest.Name,
		Email:       signUpRequest.Email,
		Password:    signUpRequest.Password,
		PhoneNumber: signUpRequest.PhoneNumber,
	}

	err := controller.UserService.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (controller *UserController) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	user, err := controller.UserService.AuthenticateUser(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	token, err := security.GenerateToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}
