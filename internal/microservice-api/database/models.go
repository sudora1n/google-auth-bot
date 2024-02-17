package orm

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Id          uint64 `json:"id" gorm:"primaryKey"`
	LanguageISO string `json:"languageiso"`

	ToTPs []ToTP `json:"totps" gorm:"foreignKey:UserId"`
}

type ToTP struct {
	gorm.Model

	Id   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`

	UserId uint64 `json:"userid"`
}
