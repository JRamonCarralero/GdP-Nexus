package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"main/config"
	"main/controllers"
	"main/models"
	"main/types"
	"main/utils"
)

// @Summary Login a user
// @Description Logs in a user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body types.LoginRequest true "Login request"
// @Success 200 {object} object "Return JWT token"
// @Failure 400 {object} object "Bad request"
// @Failure 401 {object} object "Invalid credentials"
// @Failure 500 {object} object "Internal server error"
// @Router /login [post]
func Login(c *gin.Context) {
	var req types.LoginRequest
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
		"token": tokenString,
	})
}

// @Summary Register a user
// @Description Registers a new user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body types.RegisterRequest true "Register request"
// @Success 201 {object} object "Return JWT token"
// @Failure 400 {object} object "Bad request"
// @Failure 409 {object} object "Email already registered"
// @Failure 500 {object} object "Internal server error"
// @Router /register [post]
func Register(c *gin.Context) {
	var req types.RegisterRequest
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
		"token": tokenString,
	})
}
