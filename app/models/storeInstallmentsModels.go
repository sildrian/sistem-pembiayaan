package models

type StoreInstallmentRequest struct {
	UserId 			int `json:"user_id"`
	FacilityLimitId int `json:"facility_limit_id"`
	Amount 			int `json:"amount"`
	Tenor 			int `json:"tenor"`
	StartDate 		string `json:"start_date"`

}

type StoreInstallmentResponse struct {
	UserId 				int `json:"user_id"`
	FacilityLimitId 	int `json:"facility_limit_id"`
	Amount 				int `json:"amount"`
	Tenor 				int `json:"tenor"`
	StartDate 			string `json:"start_date"`
	MonthlyInstallment 	int `json:"monthly_installment"`
	TotalMargin			int `json:"total_margin"`
	TotalPayment		int `json:"total_payment"`
	Schedule			[]*ScheduleDetail
}

type ScheduleDetail struct {
	DueDate 				string `json:""due_date"`
	InstallmentAmount 		int `json:"installment_amount"`
}
