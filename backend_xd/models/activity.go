package models

type Activity struct {
	Id           int      `gorm:"primaryKey"`
	Titulo       string   `gorm:"varchar(250);not null"`
	Dia          string   `gorm:"varchar(250);not null"`
	Horario      string   `gorm:"varchar(250);not null"`
	Imagen       string   `gorm:"varchar(250);not null"`
	Cupo         int      `gorm:"int;not null"`
	Descripcion  string   `gorm:"varchar(250);not null"`
	Categoria    Category `gorm:"foreignkey:CategoriaID"`
	CategoriaID  int
	Instructor   Instructor `gorm:"foreignkey:InstructorID"`
	InstructorID int
}

// Esta función le dice a GORM cómo se llama la tabla en la base de datos
func (Activity) TableName() string {
	return "actividades"
}
