// internal/models/order.go
package models

import (
	"gorm.io/gorm"
)

type Order struct {
	Contratto          string `gorm:"primaryKey" json:"contratto"`
	Committente        string `gorm:"not null" json:"committente"`
	DataCreazione      string `gorm:"not null" json:"data_creazione"`
	DataInizioValidita string `gorm:"not null" json:"data_inizio_validita"`
	DataFineValidita   string `gorm:"not null" json:"data_fine_validita"`
	CodiciChiave       string `gorm:"not null" json:"codice_chiave"`
	MatricolaCreazione string `gorm:"not null" json:"matricola_creazione"`
	CodiceSocieta      string `gorm:"foreignKey" json:"codice_societa"`
}

func (o *Order) Create(db *gorm.DB) error {
	return db.Create(&o).Error
}

func MigrateOrder(db *gorm.DB) error {
	return db.AutoMigrate(&Order{})
}
