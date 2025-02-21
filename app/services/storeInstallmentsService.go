package services

import (
	"database/sql"
	"sistem-pembiayaan/config"
	"sistem-pembiayaan/app/models"
	"strings"
	"strconv"
)

type CalculatoStorerService interface {
	StoreInstallmentsService(userId int, facilityLimitId int,amount int,tenorr int,startDate string) (*models.StoreInstallmentResponse, error)
}

func NewCalculatorStoreService() CalculatoStorerService {
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

func (s *calculatorService) StoreInstallmentsService(userId int, facilityLimitId int,amount int,tenorr int,startDate string) (*models.StoreInstallmentResponse, error) {
	config.InitDB()
	defer config.DB.Close()
	
	var results *models.StoreInstallmentResponse
	dynamicMarginRate := s.marginRate

	// Check if user already exists
    var exists bool
	var query string
    err := config.DB.QueryRow("SELECT COUNT(1) FROM tbtpns.user_facility WHERE user_id = @UserId", sql.Named("UserId", userId)).Scan(&exists)
    if err != nil {
        return nil, err
    }

	for _, tenor := range s.tenors {
		if tenor > 6 {
			dynamicMarginRate += 0.10
		}
		if tenorr == tenor{
			totalMargin := int((float64(amount) * dynamicMarginRate * float64(tenor)) / float64(tenor))
			totalPayment := amount + totalMargin
			monthlyInstallment := totalPayment / tenor

			if exists {
				// Update existing record
				query = `
					UPDATE tbtpns.user_facility 
					SET facility_limit_id = @FacilityLimitId,
						amount = @Amount,
						tenor = @Tenor,
						start_date = @StartDate,
						monthly_installment = @MonthlyInstallment,
						total_margin = @TotalMargin,
						total_payment = @TotalPayment
					WHERE user_id = @UserId`
			} else {
				// Insert new record
				query = `
					INSERT INTO tbtpns.user_facility 
					(user_id, facility_limit_id, amount, tenor, start_date, monthly_installment, total_margin, total_payment)
					VALUES 
					(@UserId, @FacilityLimitId, @Amount, @Tenor, @StartDate, @MonthlyInstallment, @TotalMargin, @TotalPayment)`
			}

			// Execute the query with parameters
			_, err = config.DB.Exec(query,
				sql.Named("UserId", userId),
				sql.Named("FacilityLimitId", facilityLimitId),
				sql.Named("Amount", amount),
				sql.Named("Tenor", tenorr),
				sql.Named("StartDate", startDate),
				sql.Named("MonthlyInstallment", monthlyInstallment),
				sql.Named("TotalMargin", totalMargin),
				sql.Named("TotalPayment", totalPayment),
			)
			if err != nil {
				return nil, err
			}

			var userFacilityId int
			err := config.DB.QueryRow("SELECT user_facility_id FROM tbtpns.user_facility WHERE user_id = @UserId", sql.Named("UserId", userId)).Scan(&userFacilityId)
			if err != nil {
				return nil, err
			}

			result := &models.StoreInstallmentResponse{
				UserId: 			userId,
				FacilityLimitId:	facilityLimitId,
				Amount:				amount,
				Tenor:              tenor,
				StartDate:			startDate,
				MonthlyInstallment: monthlyInstallment,
				TotalMargin:        totalMargin,
				TotalPayment:       totalPayment,
				Schedule:			ScheduleDetailService(startDate,monthlyInstallment,userFacilityId,tenor),
			}
			results = result
		}
	}

	return results, nil
}
