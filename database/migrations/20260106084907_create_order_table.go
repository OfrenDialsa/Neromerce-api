package migrations

import (
	"github.com/ofrendialsa/neromerce/database"
	"github.com/ofrendialsa/neromerce/database/entities"
	"gorm.io/gorm"
)

func init() {
	database.RegisterMigration("20260106084907_create_order_table", UpCreateOrderTable, DownCreateOrderTable)
}

func UpCreateOrderTable(db *gorm.DB) error {
	return db.AutoMigrate(&entities.Order{})
}

func DownCreateOrderTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&entities.Order{})
}
