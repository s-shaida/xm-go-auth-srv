package models

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"size:255;not null;unique"`
	Password string `json:"password" gorm:"size:255;not null;"`
}
