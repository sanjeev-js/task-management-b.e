package migrations

import (
	"fmt"
	"task-management/config"
	"task-management/models"
)

func Migrate() {
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Migration failed:", err)
	} else {
		fmt.Println("Migration completed successfully!")
	}
}
