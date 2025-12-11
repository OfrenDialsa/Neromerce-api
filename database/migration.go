package database

import (
	"github.com/ofrendialsa/neromerce/database/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entities.Migration{},
		&entities.User{},
		&entities.RefreshToken{},
		&entities.Products{},
		&entities.Category{},
	); err != nil {
		return err
	}

	manager := NewMigrationManager(db)
	return manager.Run()
}
