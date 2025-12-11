package migrations

import (
	"github.com/ofrendialsa/neromerce/database"
	"github.com/ofrendialsa/neromerce/database/entities"
	"gorm.io/gorm"
)

func init() {
	database.RegisterMigration("20251211142647_create_category_table", UpCreateCategoryTable, DownCreateCategoryTable)
}

func UpCreateCategoryTable(db *gorm.DB) error {
	return db.AutoMigrate(&entities.Category{})
}

func DownCreateCategoryTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&entities.Category{})
}
