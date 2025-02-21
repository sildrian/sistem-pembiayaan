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

var calculatorService = services.NewCalculatorService()

func TestCalculateInstallments(t *testing.T) {
	tests := []struct {
		name           string
		request        models.InstallmentRequest
		expectedStatus int
		amount         int
		marginRate     float64
		expectedResult []*models.InstallmentResponse
	}{
		{
			name:       "successful calculation",
			expectedStatus: http.StatusOK,
			amount:     10000000,
			marginRate: 0.10,
			expectedResult: []*models.InstallmentResponse{
				{
					Tenor: 6,
					MonthlyInstallment: 1833333,
					TotalMargin: 1000000,
					TotalPayment: 11000000,
				},
				{
					Tenor: 12,
					MonthlyInstallment: 1000000,
					TotalMargin: 2000000,
					TotalPayment: 12000000,
				},
				{
					Tenor: 18,
					MonthlyInstallment: 722222,
					TotalMargin: 3000000,
					TotalPayment: 13000000,
				},
				{
					Tenor: 24,
					MonthlyInstallment: 583333,
					TotalMargin: 4000000,
					TotalPayment: 14000000,
				},
				{
					Tenor: 30,
					MonthlyInstallment: 500000,
					TotalMargin: 5000000,
					TotalPayment: 15000000,
				},
				{
					Tenor: 36,
					MonthlyInstallment: 444444,
					TotalMargin: 6000000,
					TotalPayment: 16000000,
				},
			},
		},
		{
			name: "invalid amount",
			expectedStatus: http.StatusBadRequest,
			amount: 0,
			marginRate: 0.10,
			expectedResult: []*models.InstallmentResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.request)
			req := httptest.NewRequest("POST", "/calculate-installments", bytes.NewBuffer(reqBody))
			w := httptest.NewRecorder()

			controllers.CalculatorInstallments(w, req)

			if tt.expectedStatus == http.StatusBadRequest {
				assert.Equal(t, len(tt.expectedResult), 0)
			}else{
				result, err := calculatorService.CalculatorInstallments(tt.amount)
				assert.NoError(t, err)
				assert.Equal(t, len(tt.expectedResult), len(result))
				for i, expected := range tt.expectedResult {
					assert.Equal(t, expected.Tenor, result[i].Tenor)
					assert.Equal(t, expected.MonthlyInstallment, result[i].MonthlyInstallment)
					assert.Equal(t, expected.TotalMargin, result[i].TotalMargin)
					assert.Equal(t, expected.TotalPayment, result[i].TotalPayment)
				}
			}
		})
	}
}
