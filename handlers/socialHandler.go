package handlers

import (
	"encoding/json"
	"final_project/database"
	"final_project/dto"
	"final_project/models"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func IndexSocial(c *gin.Context) {
	var socials []models.Social
	db := database.GetDB()

	err := db.Preload("User").Find(&socials).Preload("User").Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "no Social available",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"result":  "sucessfully retreived Socials",
		"Socials": dto.MapSocial(socials),
	})
}

func ShowSocial(c *gin.Context) {
	var social models.Social
	session := sessions.Default(c)
	db := database.GetDB()
	id := c.Param("Social")

	err := db.Where("id = ?", id).First(&social).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "social not found",
		})
		return
	}

	fmt.Println(social)

	if session.Get("currentUser").(int) != social.User.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"social": social,
	})
}

func CreateSocial(c *gin.Context) {
	var social models.Social
	session := sessions.Default(c)
	db := database.GetDB()

	c.ShouldBind(&social)
	err := validation.ValidateStruct(
		&social,
		validation.Field(&social.Name, validation.Required),
		validation.Field(&social.SocialURL, validation.Required, is.URL),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "validation error",
			"err":     err,
		})
		return
	}

	social.User.ID = session.Get("currentUser").(int)

	err = db.Create(&social).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result":  "error uploading social",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"result":      "social uploaded",
		"id":          social.ID,
		"name":        social.Name,
		"socials_url": social.SocialURL,
		"created_at":  social.CreatedAt,
	})
}

func SocialUpdate(c *gin.Context) {
	var social models.Social
	var newSocial models.Social
	var jsonData map[string]interface{}
	session := sessions.Default(c)
	db := database.GetDB()
	id := c.Param("social")

	err := db.First(&social, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "data not found",
		})
	}

	if session.Get("currentUser").(int) != social.User.ID {
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

	newSocial.Name = jsonData["name"].(string)
	newSocial.SocialURL = jsonData["social_url"].(string)

	err = db.Model(&social).Updates(newSocial).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  "error updating Social",
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result":     "Social successfulley updated",
			"id":         social.ID,
			"name":       social.Name,
			"social_url": social.SocialURL,
			"user_id":    social.UserID,
			"updated_at": social.UpdatedAt,
		})
	}
}

func SocialDelete(c *gin.Context) {
	var social models.Social
	session := sessions.Default(c)
	db := database.GetDB()
	id := c.Param("social")

	err := db.First(&social, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "data not found",
		})
	}

	if session.Get("currentUser").(int) != social.User.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "unauthorized",
		})
		return
	}

	err = db.Delete(&social).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  "error deleting social",
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "social successfully deleted",
		})
	}
}
