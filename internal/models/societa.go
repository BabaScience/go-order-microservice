// internal/models/societa.go
package models

import (
	"gorm.io/gorm"
)

type Societa struct {
	Codice string `gorm:"primaryKey" json:"codice"`
}

func (s *Societa) GetByCodice(db *gorm.DB, code string) error {
	return db.Where("codice = ?", code).First(&s).Error
}

func (s *Societa) Create(db *gorm.DB) error {
	return db.Create(&s).Error
}

// MigrateSocieta automates the migration of the Societa model
func MigrateSocieta(db *gorm.DB) error {
	return db.AutoMigrate(&Societa{})
}
