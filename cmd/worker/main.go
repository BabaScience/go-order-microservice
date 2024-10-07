// cmd/worker/main.go
package main

import (
	"fmt"
	"goorder/internal/db"
	"goorder/internal/models"
	"goorder/internal/xmlparser"

	"github.com/joho/godotenv"
	logrus "github.com/sirupsen/logrus"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("No .env file found")
	}

	// Initialize the database
	// err = db.InitDB()
	// if err != nil {
	// 	logrus.Fatalf("Database connection error: %v", err)
	// }

	// Parse the XML file
	orders, err := xmlparser.ParseOrders("ContrattiAttivi.xml")
	if err != nil {
		logrus.Fatalf("Failed to parse XML: %v", err)
	}
	fmt.Println("Orders", orders)

	// Process each order
	for _, order := range orders {
		fmt.Printf("Processing order %s...\n", order.Contratto)
		err := processOrder(order)
		if err != nil {
			logrus.Errorf("Error processing order %s: %v", order.Contratto, err)
		} else {
			logrus.Infof("Successfully processed order %s", order.Contratto)
		}
	}

	fmt.Println("Orders processed successfully!")
}

func processOrder(order xmlparser.Order) error {
	// Initialize database if not already done
	if db.DB == nil {
		err := db.InitDB()
		if err != nil {
			return err
		}
	}

	// Start a new DB session
	database := db.DB

	// Check or create Societa
	var societa models.Societa
	err := database.First(&societa, order.CodiceSocieta).Error
	if err != nil {
		// Societa doesn't exist, create it
		societa = models.Societa{
			Codice: order.CodiceSocieta,
		}
		err = database.Create(&societa).Error
		if err != nil {
			return err
		}
	}

	// Create Posizioni
	var newPosizioni []models.Posizione
	for _, posizione := range order.Posizioni {
		newPosizione := models.Posizione{
			PosizioneContratto: posizione.PosizioneContratto,
			InizioCoperatura:   posizione.InizioCoperatura,
			FineCoperatura:     posizione.FineCoperatura,
			Quantita:           posizione.Quantita,
			UdM:                posizione.UdM,
			Divisione:          posizione.Divisione,
			Magazzino:          posizione.Magazzino,
			DilazionePagamento: posizione.DilazionePagamento,
			Listino:            posizione.Listino,
			Valuta:             posizione.Valuta,
			Materiale:          posizione.Materiale,
			ContrattoOrder:     order.Contratto,
		}

		err = newPosizione.Create(database)
		if err != nil {
			return err
		}

		newPosizioni = append(newPosizioni, newPosizione)
	}

	// Create Order
	newOrder := models.Order{
		Contratto:          order.Contratto,
		Committente:        order.Committente,
		DataCreazione:      order.DataCreazione,
		DataInizioValidita: order.DataInizioValidita,
		DataFineValidita:   order.DataFineValidita,
		CodiciChiave:       order.CodiciChiave,
		MatricolaCreazione: order.MatricolaCreazione,
		CodiceSocieta:      order.CodiceSocieta,
	}

	err = newOrder.Create(database)
	if err != nil {
		return err
	}

	return nil
}
