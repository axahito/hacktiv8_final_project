package handlers

import (
	"encoding/json"
	"final_project/database"
	"final_project/models"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IndexPhoto(c *gin.Context) {
	var photos []models.Photo
	session := sessions.Default(c)
	db := database.GetDB()

	err := db.Find(&photos).Where("user_id = ?", session.Get("currentUser")).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "no photo available",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"result": "sucessfully retreived photos",
		"photos": photos,
	})
}

func ShowPhoto(c *gin.Context) {
	var photo models.Photo
	session := sessions.Default(c)
	db := database.GetDB()
	id := c.Param("photo")

	err := db.Where("id = ?", id).First(&photo).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "photo not found",
		})
		return
	}

	fmt.Println(photo)

	if session.Get("currentUser").(int) != photo.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"photo": photo,
	})
}

func CreatePhoto(c *gin.Context) {
	var photo models.Photo
	session := sessions.Default(c)
	db := database.GetDB()
	c.ShouldBind(&photo)

	photo.UserID = session.Get("currentUser").(int)

	err := db.Create(&photo).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result":  "error uploading photo",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"result":        "photo uploaded",
		"created_photo": photo,
	})
}

func PhotoUpdate(c *gin.Context) {
	var photo models.Photo
	var newPhoto models.Photo
	var jsonData map[string]interface{}
	session := sessions.Default(c)
	db := database.GetDB()
	id := c.Param("photo")

	err := db.First(&photo, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "data not found",
		})
	}

	if session.Get("currentUser").(int) != photo.UserID {
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

	newPhoto.Title = jsonData["title"].(string)
	newPhoto.Caption = jsonData["caption"].(string)
	newPhoto.PhotoURL = jsonData["photo_url"].(string)

	err = db.Model(&photo).Updates(newPhoto).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  "error updating photo",
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result":        "photo successfulley updated",
			"updated photo": photo,
		})
	}
}

func PhotoDelete(c *gin.Context) {
	var photo models.Photo
	session := sessions.Default(c)
	db := database.GetDB()
	id := c.Param("photo")

	err := db.First(&photo, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "data not found",
		})
	}

	if session.Get("currentUser").(int) != photo.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "unauthorized",
		})
		return
	}

	err = db.Delete(&photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  "error deleting photo",
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "photo successfully deleted",
		})
	}
}
