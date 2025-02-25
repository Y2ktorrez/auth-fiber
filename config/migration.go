package config

import (
	"log"

	"github.com/mutinho/src/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		log.Fatal("Error migrating database", err)
	}
	log.Println("Database migrated")
}
