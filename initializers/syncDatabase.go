package initializers

import "jwt-go-gin/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
