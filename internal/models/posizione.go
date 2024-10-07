// internal/models/posizione.go
package models

import (
	"gorm.io/gorm"
)

type Posizione struct {
	PosizioneContratto string `gorm:"primaryKey" json:"posizione_contratto"`
	InizioCoperatura   string `gorm:"not null" json:"inizio_coperatura"`
	FineCoperatura     string `gorm:"not null" json:"fine_coperatura"`
	Quantita           string `gorm:"not null" json:"quantita"`
	UdM                string `gorm:"not null" json:"UdM"`
	Divisione          string `gorm:"not null" json:"divisione"`
	Magazzino          string `gorm:"not null" json:"magazzino"`
	DilazionePagamento string `gorm:"not null" json:"dilazione_pagamento"`
	Listino            string `gorm:"not null" json:"listino"`
	Valuta             string `gorm:"not null" json:"valuta"`
	Materiale          string `gorm:"not null" json:"materiale"`
	ContrattoOrder     string `gorm:"foreignKey" json:"contratto_order"`
}

func (d *Posizione) GetByContratto(db *gorm.DB) error {
	return db.First(&d, d.PosizioneContratto).Error
}

func (d *Posizione) Create(db *gorm.DB) error {
	return db.Create(&d).Error
}

func MigratePosizione(db *gorm.DB) error {
	return db.AutoMigrate(&Posizione{})
}
