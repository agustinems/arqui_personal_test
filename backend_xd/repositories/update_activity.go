package repositories

import (
    "proyecto2025/backend/database"
    "proyecto2025/backend/dto"
    "proyecto2025/backend/models"
)

func UpdateActividad(id int, data dto.ActividadDTO) error {
    var actividad models.Activity
    if err := database.DB.First(&actividad, id).Error; err != nil {
        return err
    }

    actividad.Titulo = data.Titulo
    actividad.Dia = data.Dia
    actividad.Horario = data.Horario
    actividad.Imagen = data.Imagen
    actividad.Cupo = data.Cupo
    actividad.Descripcion = data.Descripcion

    return database.DB.Save(&actividad).Error
}
