package models

type InstallmentRequest struct {
	Amount int `json:"amount"`
}

type InstallmentResponse struct {
	Tenor              int `json:"tenor"`
	MonthlyInstallment int `json:"monthly_installment"`
	TotalMargin        int `json:"total_margin"`
	TotalPayment       int `json:"total_payment"`
}
