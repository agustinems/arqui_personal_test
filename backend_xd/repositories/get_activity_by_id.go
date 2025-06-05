package repositories

import (
	"proyecto2025/backend/database"
	"proyecto2025/backend/dto"
	"proyecto2025/backend/models"
)

func GetActividadByID(id int) (*dto.ActividadDTO, error) {
	var actividad models.Activity
	if err := database.DB.First(&actividad, id).Error; err != nil {
		return nil, err
	}
	result := &dto.ActividadDTO{
		ID:          uint(actividad.Id),
		Titulo:      actividad.Titulo,
		Dia:         actividad.Dia,
		Horario:     actividad.Horario,
		Imagen:      actividad.Imagen,
		Cupo:        actividad.Cupo,
		Descripcion: actividad.Descripcion,
	}
	return result, nil
}
