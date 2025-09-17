package routes

import (
	"context"
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

	err := config.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo conectar a la base de datos",
		})
		return
	}
	defer config.DB.Disconnect(context.TODO())

	user, err := controllers.GetUserByEmail(config.DB, req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email incorrecto",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Password incorrecto",
		})
		return
	}

	currentUser := models.PublicUser{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		NickName:  user.NickName,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login",
		"user":    currentUser,
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

	newUser := models.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		NickName:  req.NickName,
	}
	err = config.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo conectar a la base de datos",
		})
		return
	}
	defer config.DB.Disconnect(context.TODO())

	err = controllers.CreateUser(config.DB, newUser)
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

	currentUser := models.PublicUser{
		ID:        newUser.ID,
		Email:     newUser.Email,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		NickName:  newUser.NickName,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario registrado exitosamente",
		"user":    currentUser,
	})
}
