package handlers

import (
	"net/http"

	"gin-blueprint/database"
	"gin-blueprint/models"
	"gin-blueprint/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreatePostInput struct {
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Published bool   `json:"published"`
	UserID    uint   `json:"user_id" binding:"required"`
	TagIDs    []uint `json:"tag_ids"`
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	result := database.DB.Preload("User").Preload("Tags").Find(&posts)

	if result.Error != nil {
		c.Error(utils.NewInternalError("Failed to fetch posts"))
		return
	}

	utils.SuccessResponse(c, http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	result := database.DB.
		Preload("User").
		Preload("Tags").
		First(&post, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.Error(utils.NewNotFoundError("Post not found"))
		} else {
			c.Error(utils.NewInternalError("Failed to fetch post"))
		}
		return
	}

	utils.SuccessResponse(c, http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	var input CreatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		details := make(map[string]interface{})
		details["validation_errors"] = err.Error()
		c.Error(utils.NewValidationError("Invalid input data", details))
		return
	}

	var user models.User
	if err := database.DB.First(&user, input.UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Error(utils.NewNotFoundError("User not found"))
		} else {
			c.Error(utils.NewInternalError("Failed to verify user"))
		}
		return
	}

	post := models.Post{
		Title:     input.Title,
		Content:   input.Content,
		Published: input.Published,
		UserID:    input.UserID,
	}

	if err := database.DB.Create(&post).Error; err != nil {
		c.Error(utils.NewInternalError("Failed to create post"))
		return
	}

	if len(input.TagIDs) > 0 {
		var tags []models.Tag
		database.DB.Find(&tags, input.TagIDs)
		database.DB.Model(&post).Association("Tags").Append(&tags)
	}

	database.DB.Preload("User").Preload("Tags").First(&post, post.ID)

	utils.SuccessWithMessage(c, http.StatusCreated, "Post created successfully", post)
}

func GetUserPosts(c *gin.Context) {
	userID := c.Param("id")
	var posts []models.Post

	result := database.DB.
		Where("user_id = ?", userID).
		Preload("Tags").
		Find(&posts)

	if result.Error != nil {
		c.Error(utils.NewInternalError("Failed to fetch user posts"))
		return
	}

	utils.SuccessResponse(c, http.StatusOK, posts)
}
