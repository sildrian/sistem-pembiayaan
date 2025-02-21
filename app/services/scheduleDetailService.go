package services

import (
	"database/sql"
	"sistem-pembiayaan/config"
	"sistem-pembiayaan/app/models"
	"time"
	"fmt"
)

func ScheduleDetailService (startDt string, installmentMonthly int, userFacilityId int, tenor int) ([]*models.ScheduleDetail){
	config.InitDB()
	defer config.DB.Close()

	// Check if user already exists
    var count int
	var query string
    err := config.DB.QueryRow("SELECT COUNT(detail_id) FROM tbtpns.user_facility_details WHERE user_facility_id = @UserFacilityId", sql.Named("UserFacilityId", userFacilityId)).Scan(&count)
	fmt.Println(err)
    if err != nil {
        return nil
    }

	if count > 0 && count != tenor {
		// Update existing record
		query = `
			delete from tbtpns.user_facility_details 
			WHERE user_facility_id = @UserFacilityId`
		// Execute the query with parameters
		_, err = config.DB.Exec(query,
			sql.Named("UserFacilityId", userFacilityId),
		)
		if err != nil {
			return nil
		}
	}

	var results []*models.ScheduleDetail

	// Parse the start date
    startDate, _ := time.Parse("2006-01-02", startDt)
    dueDates := generateDueDates(startDate, tenor)

	if count != tenor {
		for _, dueDate := range dueDates {
			// Insert new record
			query = `
				INSERT INTO tbtpns.user_facility_details 
				(user_facility_id, due_date, installment_amount)
				VALUES 
				(@UserFacilityId, @DueDate, @InstallmentAmount)`

			// Execute the query with parameters
			_, err = config.DB.Exec(query,
				sql.Named("UserFacilityId", userFacilityId),
				sql.Named("DueDate", dueDate),
				sql.Named("InstallmentAmount", installmentMonthly),
			)
			if err != nil {
				return nil
			}

			result := &models.ScheduleDetail{
				DueDate:              	dueDate.Format("2006-01-02"),
				InstallmentAmount: 		installmentMonthly,
			}
			results = append(results, result)
		}
	}

	return results
}

func generateDueDates(startDate time.Time, termMonths int) []time.Time {
    dueDates := make([]time.Time, termMonths)
    firstDueDate := startDate.AddDate(0, 1, 0)
    
    for i := 0; i < termMonths; i++ {
        dueDate := firstDueDate.AddDate(0, i, 0)
        dueDates[i] = dueDate
    }
    
    return dueDates
}
