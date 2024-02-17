package orm

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/sudora1n/google-auth-bot/internal/microservice-api/config"
)

func InitOrm() (*ORMFunctions, error) {
	config := config.Config
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		config.DB_Host, config.DB_User, config.DB_Password, config.DB_Name, config.DB_Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{}, &ToTP{})

	return NewORMFunctions(db), nil
}
