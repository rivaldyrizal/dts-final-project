package Model

import (
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement:true"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
