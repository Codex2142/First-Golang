package controllers

import (
	"golang-crud-login/config"
	"golang-crud-login/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user := models.User{Username: username, Password: string(hash)}

	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Sukses registrasi"})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User
	config.DB.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	session := sessions.Default(c)
	session.Set("user", user.Username)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Login success"})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Logout success"})
}
