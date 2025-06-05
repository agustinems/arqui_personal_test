package activities

import (
    "net/http"
    "strconv"

    "proyecto2025/backend/dto"
    "proyecto2025/backend/repositories"

    "github.com/gin-gonic/gin"
)

func UpdateActivity(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    var input dto.ActividadDTO
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    if err := repositories.UpdateActividad(id, input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar"})
        return
    }

    c.Status(http.StatusNoContent)
}
