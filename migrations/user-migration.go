package migrations

import (
	"fmt"
	"task-management/config"
	"task-management/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Migrate() {
	var user models.User

	if err := config.DB.Where("role = ?", "admin").First(&user).Error; err != nil {
		fmt.Println("err", err)
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Admin is not created")

			adminData := models.User{
				ID:        uuid.New().String(),
				FirstName: "William",
				LastName:  "Copper",
				Username:  "admin",
				Email:     "admin@gmail.com",
				Phone:     "9876543210",
				Password:  "Admin@123",
				Role:      "admin",
			}

			if err := config.DB.Create(&adminData).Error; err != nil {
				fmt.Printf("Error: %v\n", err.Error())
				return
			}

		} else {
			fmt.Println("Error in checking role type")
		}
	} else {
		fmt.Println("in else conditions")
	}

	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Migration failed:", err)
	} else {
		fmt.Println("Migration completed successfully!")
	}
}
