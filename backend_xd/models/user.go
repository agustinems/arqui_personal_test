package models

type User struct {
	Id           int    `gorm:"primaryKey"`
	Username     string `gorm:"type:varchar(250);not null"`
	PasswordHash string `gorm:"type:varchar(500);not null"`
	Email        string `gorm:"type:varchar(250);not null"`
	IsAdmin      bool   `gorm:"column:es_admin;type:tinyint(1);not null"`
}

func (User) TableName() string {
	return "usuarios"
}

type Users []User
