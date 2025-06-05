package activities

import (
	"net/http"
	"strconv"

	"proyecto2025/backend/repositories"

	"github.com/gin-gonic/gin"
)

func GetActivityByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	actividad, err := repositories.GetActividadByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}

	c.JSON(http.StatusOK, actividad)
}
