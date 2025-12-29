package entities

import (
	"github.com/google/uuid"
	"github.com/ofrendialsa/neromerce/pkg/helpers"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	Stock       int       `gorm:"not null;default:0" json:"stock"`
	ImageURL    string    `gorm:"type:text" json:"image_url"`

	CategoryID uint     `gorm:"not null" json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID;references:ID" json:"category"`

	Timestamp
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil && helpers.IsSQLite(tx) {
		p.ID = uuid.New()
	}
	return
}
