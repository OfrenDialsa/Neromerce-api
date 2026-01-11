package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Quantity int16     `gorm:"not null" json:"qty"`
	Price    float64   `gorm:"not null" json:"price"`

	ProductID uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`

	Timestamp
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID == uuid.Nil {
		o.ID = uuid.New()
	}
	return nil
}
