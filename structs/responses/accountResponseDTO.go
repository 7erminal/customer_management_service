package responses

import "customer_management_service/models"

type AccountDTO struct {
	StatusCode int
	Account    *models.Accounts
	StatusDesc string
}
