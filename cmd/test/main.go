package main

import (
	"fmt"
	"time"

	"goorder/internal/xmlparser"
)

func main() {
	fmt.Println("Hello, World!")

	orders, err := xmlparser.ParseOrders("ContrattiAttivi.xml")
	if err != nil {
		fmt.Println(err)
	}
	var datesString []string = []string{
		"20240601", "20240601", "20240528",
		"20250531", "20250531", "20250531",
	}
	for _, dateString := range datesString {
		fmt.Println("string: ", dateString)
		date, err := time.Parse("20060102", dateString)
		if err != nil {
			fmt.Println("Error parsing date: ", err)
		}
		dateOnly, err := time.Parse("2006-01-02", dateString)
		if err != nil {
			fmt.Println("Error parsing date: ", err)
		}
		fmt.Println("Data: ", date)
		fmt.Println("Data Only: ", dateOnly)
	}
	for _, order := range orders {
		for _, posizione := range order.Posizioni {
			fmt.Println("Pos. Contratto: ", posizione.PosizioneContratto)
			// fmt.Println("Inizio Coperatura: ", posizione.InizioCoperatura)
			// fmt.Println("Fine Coperatura: ", posizione.FineCoperatura)
			// fmt.Println("Quantita: ", posizione.Quantita)
			// fmt.Println("UdM: ", posizione.UdM)
			// fmt.Println("Divisione: ", posizione.Divisione)
			// fmt.Println("Magazzino: ", posizione.Magazzino)
			// fmt.Println("Dilazione Pagamento: ", posizione.DilazionePagamento)
			// fmt.Println("Listino: ", posizione.Listino)
			// fmt.Println("Valuta: ", posizione.Valuta)
			// fmt.Println("Materiale: ", posizione.Materiale)
			// fmt.Println("Calendario Fatturazione: ", posizione.CalendarioFatturazine)
		}
	}
}
