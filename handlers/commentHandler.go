package handlers

import (
	"encoding/json"
	"final_project/database"
	"final_project/dto"
	"final_project/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

func IndexComment(c *gin.Context) {
	var comments []models.Comment
	db := database.GetDB()

	err := db.Preload("User").Find(&comments).Preload("User").Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "no Comment available",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"result":   "sucessfully retreived comments",
		"comments": dto.MapComment(comments),
	})
}

func ShowComment(c *gin.Context) {
	var comment models.Comment
	session := sessions.Default(c)
	db := database.GetDB()
	id := c.Param("comment")

	err := db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "Comment not found",
		})
		return
	}

	if session.Get("currentUser").(int) != comment.User.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comment": comment,
	})
}

func CreateComment(c *gin.Context) {
	var comment models.Comment
	session := sessions.Default(c)
	db := database.GetDB()

	c.ShouldBind(&comment)
	err := validation.ValidateStruct(
		&comment,
		validation.Field(&comment.Message, validation.Required),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "validation error",
			"err":     err,
		})
		return
	}

	comment.User.ID = session.Get("currentUser").(int)

	err = db.Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result":  "error uploading comment",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"result":     "comment uploaded",
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	})
}

func CommentUpdate(c *gin.Context) {
	var comment models.Comment
	var newComment models.Comment
	var jsonData map[string]interface{}
	session := sessions.Default(c)
	db := database.GetDB()
	id := c.Param("comment")

	err := db.First(&comment, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "data not found",
		})
	}

	if session.Get("currentUser").(int) != comment.User.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "unauthorized",
		})
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result":  "error reading request body",
			"message": err,
		})
		return
	}

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result":  "error parsing json",
			"message": err,
		})
		return
	}

	newComment.Message = jsonData["message"].(string)

	err = db.Model(&comment).Updates(newComment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  "error updating comment",
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result":     "comment successfulley updated",
			"id":         comment.ID,
			"message":    comment.Message,
			"photo_id":   comment.PhotoID,
			"user_id":    comment.UserID,
			"updated_at": comment.UpdatedAt,
		})
	}
}

func CommentDelete(c *gin.Context) {
	var comment models.Comment
	session := sessions.Default(c)
	db := database.GetDB()
	id := c.Param("comment")

	err := db.First(&comment, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "data not found",
		})
	}

	if session.Get("currentUser").(int) != comment.User.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "unauthorized",
		})
		return
	}

	err = db.Delete(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  "error deleting comment",
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "comment successfully deleted",
		})
	}
}
