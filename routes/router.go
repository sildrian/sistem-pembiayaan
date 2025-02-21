package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"sistem-pembiayaan/app/controllers"
)

func Router() {
	router := mux.NewRouter()
	r := router.PathPrefix("/v1").Subrouter()
	r.HandleFunc("/calculate-installments", controllers.CalculatorInstallments).Methods("POST")

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
