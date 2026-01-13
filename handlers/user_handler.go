package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"gin-blueprint/database"
	"gin-blueprint/models"
	"gin-blueprint/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateUserInput struct {
	Username  string `json:"username" binding:"required,username,min=3,max=50"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,strongpassword"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateUserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"omitempty,email"`
}

func GetAllUsers(c *gin.Context) {
	var users []models.User

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	offset := (page - 1) * pageSize

	var total int64
	database.DB.Model(&models.User{}).Count(&total)

	result := database.DB.
		Limit(pageSize).
		Offset(offset).
		Find(&users)

	if result.Error != nil {
		c.Error(utils.NewInternalError("Failed to fetch users"))
		return
	}

	utils.PaginatedSuccessResponse(c, users, page, pageSize, total)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := database.DB.Preload("Posts").First(&user, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.Error(utils.NewNotFoundError("User not found"))
		} else {
			c.Error(utils.NewInternalError("Failed to fetch user"))
		}
		return
	}

	utils.SuccessResponse(c, http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		details := make(map[string]interface{})
		details["validation_errors"] = err.Error()
		c.Error(utils.NewValidationError("Invalid input data", details))
		return
	}

	user := models.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.Error(utils.NewDuplicateError("Username or email already exists"))
		} else {
			c.Error(utils.NewInternalError("Failed to create user"))
		}
		return
	}

	utils.SuccessWithMessage(c, http.StatusCreated, "User created successfully", user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.Error(utils.NewNotFoundError("User not found"))
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(utils.NewValidationError("Invalid input data", nil))
		return
	}

	if err := database.DB.Model(&user).Updates(input).Error; err != nil {
		c.Error(utils.NewInternalError("Failed to update user"))
		return
	}

	utils.SuccessWithMessage(c, http.StatusOK, "User updated successfully", user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	result := database.DB.Delete(&models.User{}, id)

	if result.Error != nil {
		c.Error(utils.NewInternalError("Failed to delete user"))
		return
	}

	if result.RowsAffected == 0 {
		c.Error(utils.NewNotFoundError("User not found"))
		return
	}

	utils.SuccessWithMessage(c, http.StatusOK, "User deleted successfully", nil)
}
