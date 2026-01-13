package database

import (
	"gin-blueprint/models"
	"log"
)

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Tag{},
	)

	if err != nil {
		log.Fatal("Migration başarısız:", err)
	}

	log.Println("✅ Migration tamamlandı!")
}
