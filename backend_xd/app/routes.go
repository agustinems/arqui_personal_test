// routes.go
package app

import (
	"proyecto2025/backend/handlers"
	"proyecto2025/backend/handlers/activities"
)

func MapRoutes() {
	// Rutas públicas
	Router.POST("/login", handlers.LoginHandler)
	Router.GET("/actividades", activities.GetAllActivities)
	// Router.GET("/actividad/:id", handlers.ObtenerActividad)

	// // Rutas de inscripción (se asume usuario logueado, pero sin middleware por ahora)
	// Router.POST("/inscribirse/:id", handlers.InscribirseActividad)
	// Router.GET("/mis-actividades", handlers.MisActividades)

	// // Rutas administrativas (sin middleware aún)
	// Router.POST("/admin/actividad", handlers.CrearActividad)
	// Router.PUT("/admin/actividad/:id", handlers.EditarActividad)
	// Router.DELETE("/admin/actividad/:id", handlers.EliminarActividad)
}
