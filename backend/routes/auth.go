package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"main/config"
	"main/controllers"
	"main/models"
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

	newUser := models.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		NickName:  req.NickName,
	}
	client, err := config.ConnectDB()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo conectar a la base de datos",
		})
		return
	}
	defer client.Disconnect(context.TODO())

	err = controllers.CreateUser(client, newUser)
	if err != nil {
		if err.Error() == "el email ya está registrado" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "El email ya está registrado",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error al registrar el usuario",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario registrado exitosamente",
		"user":    newUser,
	})
}
