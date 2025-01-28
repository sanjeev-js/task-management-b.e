package models

type User struct {
	ID        string `gorm:"primaryKey"`
	FirstName string `gorm:"not null" validate:"required,min=6,max=30"`
	LastName  string `gorm:"not null" validate:"required,min=6,max=30"`
	Username  string `gorm:"unique;not null" validate:"required,min=6,max=30"`
	Email     string `gorm:"unique;not null" validate:"required,email"`
	Phone     string `gorm:"unique;not null" validate:"required,min=10,max=12"`
	Password  string `gorm:"not null" validate:"required,min=6"`
	Role      string `gorm:"not null" validate:"required"`
	image     string `gorm:"not null" validate:"required"`
}
