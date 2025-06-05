package repositories

import (
	"log"
	"proyecto2025/backend/database"
	"proyecto2025/backend/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	log.Println("DEBUG:", user.Email, "| isAdmin:", user.IsAdmin)

	return &user, nil
}
