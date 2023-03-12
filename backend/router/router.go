package router

import (
	"go-postgres-crud/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/gas", controller.AmbilSemuaData).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/gas/{id}", controller.AmbilData).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/gas", controller.TmbhData).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/gas/{id}", controller.UpdateData).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/gas/{id}", controller.HapusData).Methods("DELETE", "OPTIONS")

	return router
}
