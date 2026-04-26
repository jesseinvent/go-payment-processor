package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jesseinvent/go-payment-processor/internal/pkg/response"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) UserHandler {
	return UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	
  var createUserDto CreateUserDto

  err := c.ShouldBindJSON(&createUserDto)

  if err != nil {
	c.JSON(400, response.Error(fmt.Sprintf("Invalid request format - %s", err.Error())))

	return
  }

  user, err := h.userService.Create(createUserDto.Email, createUserDto.PhoneNumber, createUserDto.Name)

  if err != nil {
	log.Print(err.Error())

	c.JSON(http.StatusBadRequest, response.Error("Could not create user."))

	return
  }

  createUserResponse := UserResponse{
	ID: user.ID,
	Email: user.Email,
	PhoneNumber: user.PhoneNumber,
	Name: user.PhoneNumber,
	CreatedAt: user.CreatedAt,
  }

  c.JSON(http.StatusCreated, response.Success("User successfully created", createUserResponse))
 }

func (h *UserHandler) GetUserById(c *gin.Context) {
	
	userId := c.Param("userId");

	if userId == "" {
		c.JSON(http.StatusBadRequest, response.Error("Please provide user Id"))
		return
	}

	id, err := strconv.Atoi(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid userId"))
		return
	}

	user, err := h.userService.GetByID(id)

	if user == nil {
		c.JSON(http.StatusBadRequest, response.Error("User not found"))
		return
	}

	userResponse := UserResponse{
		ID: user.ID,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
		Name: user.PhoneNumber,
		CreatedAt: user.CreatedAt,
  	}

	c.JSON(http.StatusOK, response.Success("User found", userResponse))
}