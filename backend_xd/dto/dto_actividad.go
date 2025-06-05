package dto

type ActividadDTO struct {
	ID          uint   `json:"id"`
	Titulo      string `json:"titulo"`
	Dia         string `json:"dia"`
	Horario     string `json:"horario"`
	Imagen      string `json:"imagen"`
	Cupo        int    `json:"cupo"`
	Descripcion string `json:"descripcion"`
}
