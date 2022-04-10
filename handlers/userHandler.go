package handlers

import (
	"encoding/json"
	"final_project/database"
	"final_project/helpers"
	"final_project/models"
	"io/ioutil"
	"net/http"

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
		"message": "successfully registered user",
		// "data":    &user,
	}

	c.JSON(http.StatusCreated, result)
}

func UserLogin(c *gin.Context) {
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

	token := helpers.GenerateToken(user.ID, user.Email)

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}
