package helpers

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func IsSQLite(db *gorm.DB) bool {
	return db.Dialector.Name() == "sqlite"
}

func GenerateID(db *gorm.DB) string {
	if IsSQLite(db) {
		return uuid.New().String()
	}
	return ""
}
