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

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  "error parsing json",
			"message": err,
		})
	}

	json.Unmarshal(body, &user)

	if !helpers.Email(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  "error creating user",
			"message": "email is invalid",
		})
	}

	db := database.GetDB()

	c.ShouldBind(&user)
	err = db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error creating user",
			"err":     err,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "successfully registered user",
		"created_user": user,
	})
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
	var jsonData map[string]interface{}
	session := sessions.Default(c)
	db := database.GetDB()
	id, err := strconv.Atoi(c.Param("user"))
	if session.Get("currentUser") != id {
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

	newUser.Age = jsonData["age"].(int)
	newUser.Email = jsonData["email"].(string)
	newUser.Username = jsonData["username"].(string)

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
	session := sessions.Default(c)
	db := database.GetDB()
	id, err := strconv.Atoi(c.Param("user"))
	if session.Get("currentUser") != id {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "unauthorized",
		})
		return
	}

	err = db.First(&user, id).Error
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
