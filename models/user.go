package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique; not null"`
	Password []byte `json:"-" gorm:"not null"`
}
