package main

import (
	"fmt"
	"log"
	"os"

	"proyecto2025/backend/app"
	"proyecto2025/backend/config"
	"proyecto2025/backend/database"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env") // ⬅️ carga explícita del archivo
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}
	fmt.Println("DEBUG DB_DSN:", os.Getenv("DB_DSN")) // mantener para verificar
}

func main() {
	// Inicializar base de datos
	database.InitDB()

	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Fatal("No se pudo obtener el manejador de conexión:", err)
	}
	defer sqlDB.Close()

	// Inicializar router y rutas
	app.InitRouter()
	app.MapRoutes()

	log.Println("Servidor escuchando en http://localhost:" + config.Port())
	app.Router.Run(":" + config.Port())
}
