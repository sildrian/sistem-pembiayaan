package controllers

import (
	"encoding/json"
	"net/http"

	"sistem-pembiayaan/app/models"
	"sistem-pembiayaan/app/services"
	"sistem-pembiayaan/app/library"
)

var calculatorStoreService = services.NewCalculatorStoreService()

func StoreInstallments(w http.ResponseWriter, r *http.Request) {
	var req models.StoreInstallmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		library.Res_400(w, "Invalid request body")
		return
	}

	// amounnt validation
	if req.Amount <= 0 {
		library.Res_400(w, "Amount must be greater than 0")
		return
	}

	// tenor validation
	checkTenor, err := services.TenorValidation(req.Tenor)
	if err != nil {
		library.Res_400(w, "please check your tenor input")
		return
	}
	if checkTenor == false {
		library.Res_400(w, "tenor not in list (6,12,18,24,30,36)")
		return
	}

	// validate date format
	if !services.ValidateDateFormat(req.StartDate) {
		library.Res_400(w, "Start_date not valid. Please input start_date like 2025-02-21")
		return
	}

	// facility limit validation
	checkUserFacilityLimit := services.UserFacilityLimit(req.FacilityLimitId, req.UserId)
	if checkUserFacilityLimit == nil {
		library.Res_400(w, "Facility Limit Id/User Id not found")
		return
	}
	if checkUserFacilityLimit.LimitAmount < req.Amount {
		library.Res_400(w, "Your Limit not enough")
		return
	}

	result, err := calculatorStoreService.StoreInstallmentsService(req.UserId,req.FacilityLimitId,req.Amount,req.Tenor,req.StartDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	library.Res_200(w, "success get data", result)
}
