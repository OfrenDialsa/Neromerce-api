package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name string    `gorm:"type:varchar(100);not null;unique" json:"name"`

	Products []Products `gorm:"foreignKey:CategoryID" json:"products"`

	Timestamp
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}
