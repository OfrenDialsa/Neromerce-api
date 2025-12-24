package entities

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(100);not null;unique" json:"name"`

	Products []Products `gorm:"foreignKey:CategoryID" json:"products"`

	Timestamp
}
