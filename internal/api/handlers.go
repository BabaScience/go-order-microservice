// internal/api/handlers.go
package api

import (
	"encoding/json"
	"goorder/internal/db"
	"goorder/internal/models"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Societa Endpoints
	router.HandleFunc("/api/v1/societa", GetSocietaListHandler).Methods("GET")
	router.HandleFunc("/api/v1/societa/{codice:[0-9]+}", GetSocietaByCodiceHandler).Methods("GET")
	router.HandleFunc("/api/v1/societa", CreateSocietaByCodiceHandler).Methods("POST")

	// Deposito Endpoints
	router.HandleFunc("/api/v1/posizione", GetPosizioneListHandler).Methods("GET")
	router.HandleFunc("/api/v1/posizione/{contratto:[0-9]+}", GetPosizioneByContrattoHandler).Methods("GET")
	router.HandleFunc("/api/v1/posizione", CreatePosizioneHandler).Methods("POST")

	// Order Endpoint
	router.HandleFunc("/api/v1/orders", GetOrdersHandler).Methods("GET")
	router.HandleFunc("/api/v1/orders", CreateOrderHandler).Methods("POST")

	return router
}

// CreateSocietaByCodiceHandler creates a new Societa by Codice
func CreateSocietaByCodiceHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Unmarshal the request body into a map
	data := map[string]string{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Create a new Societa
	var s models.Societa = models.Societa{Codice: data["codice"]}

	errSavingData := s.Create(db.DB)
	if errSavingData != nil {
		http.Error(w, errSavingData.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(s)
}

func CreateSocietaHandler(w http.ResponseWriter, r *http.Request) {
	var s models.Societa
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&s)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

// Posizione Handlers
func GetPosizioneByContrattoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contratto, _ := strconv.Atoi(vars["contratto"])

	var d models.Posizione
	result := db.DB.First(&d, contratto)
	if result.Error != nil {
		http.Error(w, "Posizione not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(d)
}

func CreatePosizioneHandler(w http.ResponseWriter, r *http.Request) {
	var d models.Posizione
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&d)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(d)
}

// Order Handler
func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var o models.Order
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&o)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(o)
}

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	result := db.DB.Find(&orders)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}

func GetSocietaListHandler(w http.ResponseWriter, r *http.Request) {
	var societas []models.Societa
	result := db.DB.Find(&societas)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(societas)
}

// GetSocietaByCodiceHandler gets a Societa by Codice
func GetSocietaByCodiceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["codice"]

	var s models.Societa
	result := db.DB.Where("codice = ?", code).First(&s)
	if result.Error != nil {
		http.Error(w, "Societa not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(s)
}

func GetPosizioneListHandler(w http.ResponseWriter, r *http.Request) {
	var posizioneList []models.Posizione

	// Check query parameters
	query := r.URL.Query()
	if len(query) > 0 {
		var posizione models.Posizione
		for key, value := range query {
			if key == "contratto_order" {
				posizione.ContrattoOrder = value[0]
			}
		}
		result := db.DB.Where(&posizione).Find(&posizioneList)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(posizioneList)
		return
	}

	result := db.DB.Find(&posizioneList)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(posizioneList)
}
