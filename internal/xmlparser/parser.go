package xmlparser

import (
	"encoding/xml"
	"os"
)

type OrdersXML struct {
	XMLName xml.Name `xml:"RicezioneContrattoAttivoConfirmation"`
	Orders  []Order  `xml:"rows>row"`
}

type Order struct {
	XMLName            xml.Name    `xml:"row"`
	Contratto          string      `xml:"contratto"`
	Committente        string      `xml:"committente"`
	DataCreazione      string      `xml:"data_creazione"`
	DataInizioValidita string      `xml:"data_inizio_validita"`
	DataFineValidita   string      `xml:"data_fine_validita"`
	CodiciChiave       string      `xml:"codici_chiave"`
	MatricolaCreazione string      `xml:"matricola_creazione"`
	CodiceSocieta      string      `xml:"codSocieta"`
	Posizioni          []Posizione `xml:"posizioni>posizione"`
}

type Posizione struct {
	XMLName               xml.Name `xml:"posizione"`
	PosizioneContratto    string   `xml:"posizione_contratto"`
	InizioCoperatura      string   `xml:"inizio_copertura"`
	FineCoperatura        string   `xml:"fine_copertura"`
	Quantita              string   `xml:"quantita"`
	UdM                   string   `xml:"UdM"`
	Divisione             string   `xml:"divisione"`
	Magazzino             string   `xml:"magazzino"`
	DilazionePagamento    string   `xml:"dilazione_pagamento"`
	Listino               string   `xml:"listino"`
	Valuta                string   `xml:"valuta"`
	CalendarioFatturazine string   `xml:"calendario_fatturazione"`
	Materiale             string   `xml:"materiale"`
}

func ParseOrders(filename string) ([]Order, error) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	var orders OrdersXML
	decoder := xml.NewDecoder(xmlFile)
	err = decoder.Decode(&orders)
	if err != nil {
		return nil, err
	}
	return orders.Orders, nil
}
