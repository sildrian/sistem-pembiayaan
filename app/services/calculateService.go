package services

import (
	"database/sql"
	"sistem-pembiayaan/config"
	"sistem-pembiayaan/app/models"
	"strings"
	"strconv"
	"fmt"
)

type CalculatorService interface {
	CalculatorInstallments(amount int) ([]*models.InstallmentResponse, error)
}

type calculatorService struct {
	marginRate float64
	tenors []int
}

func NewCalculatorService() CalculatorService {
	config.InitDB()
	defer config.DB.Close()

	var tenorString string
	var tenors []int
	err := config.DB.QueryRow("SELECT tenor_value FROM tbtpns.tenor where tenor_id=1").Scan(&tenorString)
	if err != nil {
		if err == sql.ErrNoRows {
			return &calculatorService{
				marginRate: 0.10,
				tenors: tenors,
			}
		}
		return &calculatorService{
			marginRate: 0.10,
			tenors: tenors,
		}
	}

	tenorStrings := strings.Split(tenorString, ",")
	for _, t := range tenorStrings {
		tenor, err := strconv.Atoi(strings.TrimSpace(t))
		if err != nil {
			return &calculatorService{
				marginRate: 0.10,
				tenors: tenors,
			}
		}
		tenors = append(tenors, tenor)
	}

	return &calculatorService{
		marginRate: 0.10,
		tenors: tenors,
	}
}

func (s *calculatorService) CalculatorInstallments(amount int) ([]*models.InstallmentResponse, error) {
	var results []*models.InstallmentResponse
	dynamicMarginRate := s.marginRate
	for _, tenor := range s.tenors {
		if tenor > 6 {
			dynamicMarginRate += 0.10
		}
		totalMargin := int((float64(amount) * dynamicMarginRate * float64(tenor)) / float64(tenor))
		totalPayment := amount + totalMargin
		monthlyInstallment := totalPayment / tenor

		result := &models.InstallmentResponse{
			Tenor:              tenor,
			MonthlyInstallment: monthlyInstallment,
			TotalMargin:        totalMargin,
			TotalPayment:       totalPayment,
		}

		results = append(results, result)
	}

	return results, nil
}
