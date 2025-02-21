package services

import (
	"sistem-pembiayaan/config"
	"sistem-pembiayaan/app/models"
	"strconv"
	"strings"
	"time"
)

func UserFacilityLimit (FLimitId int, Userid int) (*models.UserFacilityLimitResponse){
	config.InitDB()
	defer config.DB.Close()

	var results *models.UserFacilityLimitResponse
	var (
		column1 string
		column2 string
	)

	err := config.DB.QueryRow("SELECT facility_limit_id,limit_amount FROM tbtpns.facility_limit where facility_limit_id=@p1 and user_id=@p2", FLimitId,Userid).Scan(&column1, &column2)
	if err != nil {
		return results
	}

	// Convert string to int and handle errors
	facilityLimitId, err := strconv.Atoi(column1)
	if err != nil {
		return results
	}

	limitAmount, err := strconv.Atoi(column2)
	if err != nil {
		return results
	}

	return &models.UserFacilityLimitResponse{
		FacilityLimitId: facilityLimitId,
		LimitAmount:     limitAmount,
	}
}

func TenorValidation (TenorParam int) (bool, error){
	config.InitDB()
	defer config.DB.Close()

	var tenorString string
	err := config.DB.QueryRow("SELECT tenor_value FROM tbtpns.tenor where tenor_id=1").Scan(&tenorString)
	if err != nil {
		return false, err
	}

	tenorStrings := strings.Split(tenorString, ",")
	for _, t := range tenorStrings {
		tenor, err := strconv.Atoi(strings.TrimSpace(t))
		if err != nil {
			return false, err
		}
		if tenor == TenorParam {
			return true, nil
		}
	}

	return false, nil
}


func ValidateDateFormat(dateStr string) bool {
	// Define the expected date format
	const layout = "2006-01-02"

	// Attempt to parse the date string
	_, err := time.Parse(layout, dateStr)
	if err != nil {
		return false // Parsing failed, format is incorrect
	}

	return true // Parsing succeeded, format is correct
}
