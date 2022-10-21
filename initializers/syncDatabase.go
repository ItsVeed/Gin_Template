package initializers

import "github.com/ItsVeed/Gin_Template/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
