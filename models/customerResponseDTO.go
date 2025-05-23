package models

// import (
// 	"github.com/beego/beego/v2/client/orm"
// )

type CustomerResponseDTO struct {
	StatusCode int
	Customer   *Customers
	StatusDesc string
}

type CustomerEmergencyContactResponseDTO struct {
	StatusCode               int
	CustomerEmergencyContact *Customer_emergency_contacts
	StatusDesc               string
}

type CustomerGuarantorResponseDTO struct {
	StatusCode        int
	CustomerGuarantor *Customer_guarantors
	StatusDesc        string
}
