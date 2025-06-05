package handlers

import (
	"log"
	"net/http"

	"proyecto2025/backend/dto"
	"proyecto2025/backend/repositories"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	user, err := repositories.GetUserByEmail(req.Email)
	if err != nil {
		log.Println("No se encontró usuario:", req.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email o contraseña incorrectos"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		log.Println("Contraseña inválida")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email o contraseña incorrectos"})
		return
	}

	isAdmin := user.IsAdmin

	token, err := repositories.GenerateJWT(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar token"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{Token: token, IsAdmin: isAdmin})
}
