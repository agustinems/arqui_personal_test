package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DB_DSN") // o config.Dsn() si usás config
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error abriendo conexión:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("No se pudo obtener la conexión base:", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("No se pudo hacer ping a la base de datos:", err)
	}

	fmt.Println("Conexión exitosa a MySQL.")
}
