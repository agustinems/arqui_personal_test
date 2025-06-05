package repositories

import (
	"proyecto2025/backend/database"
	"proyecto2025/backend/dto"
	"proyecto2025/backend/models"
)

func GetAllActividades() ([]dto.ActividadDTO, error) {
	var actividades []models.Activity

	if err := database.DB.Find(&actividades).Error; err != nil {
		return nil, err
	}

	var result []dto.ActividadDTO
	for _, a := range actividades {
		result = append(result, dto.ActividadDTO{
			ID:          uint(a.Id),
			Titulo:      a.Titulo,
			Dia:         a.Dia,
			Horario:     a.Horario,
			Imagen:      a.Imagen,
			Cupo:        a.Cupo,
			Descripcion: a.Descripcion,
		})
	}

	return result, nil
}
