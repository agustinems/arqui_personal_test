package activities

import (
	"net/http"
	"proyecto2025/backend/repositories"

	"github.com/gin-gonic/gin"
)

func GetAllActivities(c *gin.Context) {
	actividades, err := repositories.GetAllActividades()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las actividades"})
		return
	}
	c.JSON(http.StatusOK, actividades)
}
