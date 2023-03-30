package controllers

import (
	"coding-test/backend/database"
	"coding-test/backend/helper"
	"coding-test/backend/models"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var db = database.Connect()

var Store = sessions.NewCookieStore([]byte("secret"))

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}

func Sortnum(c *gin.Context) {
	var numbers []int
	if err := c.BindJSON(&numbers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sort.Ints(numbers)

	c.JSON(http.StatusOK, gin.H{"numbers": numbers})
}

func Login(c *gin.Context) {
	var login models.Login
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong format"})
		return
	}

	var check models.Login
	db.Table("users").Where("email = ?", login.Email).Take(&check)
	if check.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect"})
		return
	}

	if match := helper.CheckPasswordHash(login.Password, check.Password); !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect"})
		return
	}

	JWT, _ := helper.GenerateJWT(c, login.Email)

	c.JSON(http.StatusOK, gin.H{"login": true, "access_token": JWT})
}

func Is_auth(c *gin.Context) {
	email, _ := c.Get("email")
	c.JSON(http.StatusOK, gin.H{"auth": true, "email": email})
}
