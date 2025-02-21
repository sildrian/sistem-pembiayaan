package controllers

import (
	"encoding/json"
	"net/http"

	"sistem-pembiayaan/app/models"
	"sistem-pembiayaan/app/services"
	"sistem-pembiayaan/app/library"
)

var calculatorService = services.NewCalculatorService()

func CalculatorInstallments(w http.ResponseWriter, r *http.Request) {
	var req models.InstallmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		library.Res_400(w, "Invalid request body")
		return
	}

	if req.Amount <= 0 {
		library.Res_400(w, "Amount must be greater than 0")
		return
	}

	result, err := calculatorService.CalculatorInstallments(req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	library.Res_200(w, "success get data", result)
}
