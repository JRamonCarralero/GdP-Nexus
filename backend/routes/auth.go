package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"main/config"
	"main/controllers"
	"main/models"
	struts "main/structs"
	"main/utils"
)

// Login handles a login request and returns a JWT token for the user.
// The request must contain a valid email and password.
// If the email or password are incorrect, a 401 Unauthorized status is returned.
// If an error occurs while generating the token, a 500 Internal Server Error status is returned.
// The response contains the user's information and the JWT token.
// If the request is successful, a 200 OK status is returned.
//
// @param c *gin.Context
func Login(c *gin.Context) {
	var req struts.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := controllers.GetUserByEmail(config.DB, req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Credenciales incorrectas",
		})
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Credenciales incorrectas",
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

	tokenString, err := utils.GenerateToken(currentUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al generar el token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login",
		"user":    currentUser,
		"token":   tokenString,
	})
}

// Register handles a register request and returns a JWT token for the user.
// The request must contain a valid email, password, first name, last name, and nickname.
// If the email is already registered, a 409 Conflict status is returned.
// If an error occurs while generating the token, a 500 Internal Server Error status is returned.
// The response contains the user's information and the JWT token.
// If the request is successful, a 200 OK status is returned.
//
// @param c *gin.Context
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

	err = controllers.CreateUser(config.DB, newUser)
	if err != nil {
		if err.Error() == "Email is already registered" {
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

	tokenString, err := utils.GenerateToken(currentUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al generar el token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario registrado",
		"user":    currentUser,
		"token":   tokenString,
	})
}
