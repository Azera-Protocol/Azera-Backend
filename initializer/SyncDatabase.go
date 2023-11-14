package initializer

import "azera-backend/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Customization{})
}
