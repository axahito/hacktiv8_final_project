package handlers

import (
	"encoding/json"
	"final_project/database"
	"final_project/helpers"
	"final_project/models"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var user models.User
	var result gin.H

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		result = gin.H{
			"result":  "error parsing json",
			"message": err,
		}

		c.JSON(http.StatusBadRequest, result)
	}

	json.Unmarshal(body, &user)

	if !helpers.Email(user.Email) {
		result = gin.H{
			"result":  "error creating user",
			"message": "email is invalid",
		}

		c.JSON(http.StatusBadRequest, result)
	}

	db := database.GetDB()

	c.ShouldBind(&user)
	err = db.Create(&user).Error
	if err != nil {
		result = gin.H{
			"message": "error creating user",
			"err":     err,
		}

		c.JSON(http.StatusInternalServerError, result)
	}

	result = gin.H{
		"message":      "successfully registered user",
		"created_user": user,
	}

	c.JSON(http.StatusCreated, result)
}

func UserLogin(c *gin.Context) {
	session := sessions.Default(c)
	db := database.GetDB()
	user := models.User{}
	c.ShouldBind(&user)

	password := ""
	password = user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email / password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email / password",
		})
		return
	}

	session.Set("currentUser", user.ID)
	session.Save()

	token := helpers.GenerateToken(user.ID, user.Email)

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}

func UserUpdate(c *gin.Context) {
	var user models.User
	var newUser models.User
	// var currentUser models.User
	session := sessions.Default(c)
	db := database.GetDB()
	id, err := strconv.Atoi(c.Param("user"))
	if session.Get("currentUser") != id {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "unauthorized",
		})
		return
	}

	c.Bind(&newUser)

	err = db.First(&user, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "data not found",
		})
	}

	err = db.Model(&user).Updates(newUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  "error updating user",
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result":       "user successfulley updated",
			"updated user": user,
		})
	}
}

func UserDelete(c *gin.Context) {
	var user models.User
	db := database.GetDB()
	id := c.Param("user")

	err := db.First(&user, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "data not found",
		})
	}

	err = db.Delete(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  "error deleting user",
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "user successfully deleted",
		})
	}
}
