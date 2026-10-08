package responses

import (
	"customer_management_service/models"
	"time"
)

type Identification_types struct {
	IdentificationTypeId string
	Name                 string
	Code                 string
	DateCreated          time.Time
	DateModified         time.Time
	CreatedBy            string
	ModifiedBy           string
	Active               int
}

type Branches struct {
	BranchId     string
	Branch       string
	Country      string
	Location     string
	PhoneNumber  string
	Active       int
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    string
	ModifiedBy   string
}

type ShopBranches struct {
	Id           int64
	Shop         *Shops
	Branch       *Branches
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    string
	ModifiedBy   string
	Active       int
}

type Shops struct {
	ShopId              string
	ShopName            string
	ShopDescription     string
	ShopAssistantName   string
	ShopAssistantNumber string
	PhoneNumber         string
	Email               string
	Image               string
	ShopLocation        string
	DateCreated         time.Time
	DateModified        time.Time
	CreatedBy           string
	ModifiedBy          string
	Active              int
	ShopBranches        []*ShopBranches
}

type Customer_categories struct {
	CustomerCategoryId string
	Category           string
	Code               string
	Description        string
	DateCreated        time.Time
	DateModified       time.Time
	CreatedBy          string
	ModifiedBy         string
	Active             int
}

type Customers struct {
	CustomerId           string
	CustomerNumber       string
	FullName             string
	ImagePath            string
	Email                string
	PhoneNumber          string
	Gender               string
	Location             string
	IdentificationType   *Identification_types
	IdentificationNumber string
	Branch               *Branches
	Shop                 *Shops
	CustomerCategory     *Customer_categories
	Nickname             string
	Dob                  time.Time
	DateCreated          time.Time
	DateModified         time.Time
	CreatedBy            string
	ModifiedBy           string
	Active               int
	LastTxnDate          time.Time
	EmergencyContacts    []*Customer_emergency_contacts
	Guarantors           []*Customer_guarantors
}

type Customer_emergency_contacts struct {
	CustomerEmergencyContactId int64
	Name                       string
	Contact                    string
	Customer                   *Customers
	DateCreated                time.Time
	DateModified               time.Time
	CreatedBy                  string
	ModifiedBy                 string
}

type Customer_guarantors struct {
	CustomerGuarantorId string
	Name                string
	Contact             string
	Customer            *Customers
	DateCreated         time.Time
	DateModified        time.Time
	CreatedBy           string
	ModifiedBy          string
}

type CustomersDTO struct {
	StatusCode int
	Customers  *[]models.Customers
	StatusDesc string
}
