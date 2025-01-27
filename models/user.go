package models

type User struct {
	ID       string `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null" validate:"required,email"`
	Password string `gorm:"not null" validate:"required,min=6"`
	Role     string `gorm:"not null" validate:"required"`
}
