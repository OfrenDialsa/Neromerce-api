package migrations

import (
	"github.com/ofrendialsa/neromerce/database"
	"github.com/ofrendialsa/neromerce/database/entities"
	"gorm.io/gorm"
)

func init() {
	database.RegisterMigration("20251209195027_create_product_table", UpCreateProductsTable, DownCreateProductsTable)
}

func UpCreateProductsTable(db *gorm.DB) error {
	return db.AutoMigrate(&entities.Product{})
}

func DownCreateProductsTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&entities.Product{})
}
