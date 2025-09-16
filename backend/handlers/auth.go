package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	struts "main/structs"
)

func Login(c *gin.Context) {
	var req struts.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Printf("Email recibido: %s", req.Email)
	fmt.Printf("Password recibido: %s", req.Password)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login",
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout",
	})
}

func Register(c *gin.Context) {
	var req struts.RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("\nEmail recibido: %s", req.Email)
	fmt.Printf("\nPassword recibido: %s", hashedPassword)
	fmt.Printf("\nFirstName recibido: %s", req.FirstName)
	fmt.Printf("\nLastName recibido: %s", req.LastName)
	fmt.Printf("\nNickName recibido: %s", req.NickName)
	fmt.Printf("\n")

	c.JSON(http.StatusOK, gin.H{
		"message": "Register",
	})
}
