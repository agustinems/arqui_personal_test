package dto

// Estructura para parsear el JSON de login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandler maneja las peticiones de inicio de sesión con Gin
type LoginResponse struct {
	Token   string `json:"token"`
	IsAdmin bool   `json:"isAdmin"`
}
