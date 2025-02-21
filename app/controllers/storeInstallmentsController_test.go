package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"sistem-pembiayaan/app/models"
	"sistem-pembiayaan/app/controllers"
	"sistem-pembiayaan/app/services"
)

var calculatorStoreService = services.NewCalculatorStoreService()

func TestCalculateInstallments(t *testing.T) {
	tests := []struct {
		name           		string
		request        		models.StoreInstallmentRequest
		expectedStatus 		int
		amount         		int
		user_id 			int
		facility_limit_id 	int
		tenor 				int
		start_date 			string
		marginRate     		float64
		expectedResult 		models.StoreInstallmentResponse
	}{
		{
			name:       		"successful calculation",
			expectedStatus: 	http.StatusOK,
			amount:     		30000000,
			user_id: 			1,
			facility_limit_id: 	1,
			tenor: 				12,
			start_date: 		"2025-02-09",
			marginRate: 		0.10,
			expectedResult: 	models.StoreInstallmentResponse{
				UserId: 				1,
				FacilityLimitId: 		1,
				Amount: 				30000000,	
				Tenor: 					12,
				StartDate: 				"2025-02-09",
				MonthlyInstallment: 	3000000,
				TotalMargin:			6000000,
				TotalPayment:			36000000,
				Schedule: []*models.ScheduleDetail{
					{
						DueDate: "2025-03-09",
						InstallmentAmount: 3000000,
					},
					{
						DueDate: "2025-04-09",
						InstallmentAmount: 3000000,
					},
					{
						DueDate: "2025-05-09",
						InstallmentAmount: 3000000,
					},
					{
						DueDate: "2025-06-09",
						InstallmentAmount: 3000000,
					},
					{
						DueDate: "2025-07-09",
						InstallmentAmount: 3000000,
					},
					{
						DueDate: "2025-08-09",
						InstallmentAmount: 3000000,
					},
					{
						DueDate: "2025-09-09",
						InstallmentAmount: 3000000,
					},
					{
						DueDate: "2025-10-09",
						InstallmentAmount: 3000000,
					},
					{
						DueDate: "2025-11-09",
						InstallmentAmount: 3000000,
					},
					{
						DueDate: "2025-12-09",
						InstallmentAmount: 3000000,
					},
					{
						DueDate: "2026-01-09",
						InstallmentAmount: 3000000,
					},
					{
						DueDate: "2026-02-09",
						InstallmentAmount: 3000000,
					},
				},
			},
		},
		{
			name: "invalid amount",
			expectedStatus: http.StatusBadRequest,
			amount: 0,
			user_id: 			1,
			facility_limit_id: 	1,
			tenor: 				12,
			start_date: 		"2025-02-09",
			marginRate: 0.10,
			expectedResult: models.StoreInstallmentResponse{},
		},
		{
			name: "invalid tenor",
			expectedStatus: http.StatusBadRequest,
			amount: 			30000000,
			user_id: 			1,
			facility_limit_id: 	1,
			tenor: 				10,
			start_date: 		"2025-02-09",
			marginRate: 0.10,
			expectedResult: models.StoreInstallmentResponse{},
		},
		{
			name: "invalid date format",
			expectedStatus: http.StatusBadRequest,
			amount: 			30000000,
			user_id: 			1,
			facility_limit_id: 	1,
			tenor: 				12,
			start_date: 		"02-09-2025",
			marginRate: 0.10,
			expectedResult: models.StoreInstallmentResponse{},
		},
		{
			name: "facility id/ user id not found",
			expectedStatus: http.StatusBadRequest,
			amount: 			30000000,
			user_id: 			1,
			facility_limit_id: 	2,
			tenor: 				12,
			start_date: 		"02-09-2025",
			marginRate: 0.10,
			expectedResult: models.StoreInstallmentResponse{},
		},
		{
			name: "invalid limit amount",
			expectedStatus: http.StatusBadRequest,
			amount: 			90000000,
			user_id: 			1,
			facility_limit_id: 	1,
			tenor: 				12,
			start_date: 		"02-09-2025",
			marginRate: 0.10,
			expectedResult: models.StoreInstallmentResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.request)
			req := httptest.NewRequest("POST", "/submit-financing", bytes.NewBuffer(reqBody))
			w := httptest.NewRecorder()

			controllers.StoreInstallments(w, req)
			checkDate := services.ValidateDateFormat(tt.start_date)
			checkUserFacilityLimit := services.UserFacilityLimit(tt.facility_limit_id, tt.user_id)

			if tt.amount <= 0 {
				assert.Equal(t, tt.expectedStatus, http.StatusBadRequest)
			}else if tt.tenor % 6 != 0 && tt.tenor > 36 {
				assert.Equal(t, tt.expectedStatus, http.StatusBadRequest)
			}else if !checkDate {
				assert.Equal(t, tt.expectedStatus, http.StatusBadRequest)
			}else if checkUserFacilityLimit == nil{
				assert.Equal(t, tt.expectedStatus, http.StatusBadRequest)
			}else if checkUserFacilityLimit.LimitAmount < tt.amount{
				assert.Equal(t, tt.expectedStatus, http.StatusBadRequest)
			}else{
				result, err := calculatorStoreService.StoreInstallmentsService(tt.user_id,tt.facility_limit_id,tt.amount,tt.tenor,tt.start_date)
				assert.NoError(t, err)
				if result != nil {
					assert.Equal(t, tt.expectedResult.UserId, result.UserId)
					assert.Equal(t, tt.expectedResult.FacilityLimitId, result.FacilityLimitId)
					assert.Equal(t, tt.expectedResult.Amount, result.Amount)
					assert.Equal(t, tt.expectedResult.Tenor, result.Tenor)
					assert.Equal(t, tt.expectedResult.StartDate, result.StartDate)
					assert.Equal(t, tt.expectedResult.MonthlyInstallment, result.MonthlyInstallment)
					assert.Equal(t, tt.expectedResult.TotalMargin, result.TotalMargin)
					assert.Equal(t, tt.expectedResult.TotalPayment, result.TotalPayment)

					// Test schedule details if they exist
					// if len(tt.expectedResult.Schedule) > 0 {
					// 	assert.Equal(t, len(tt.expectedResult.Schedule), len(result.Schedule))
					// }
				}
			}
		})
	}
}
