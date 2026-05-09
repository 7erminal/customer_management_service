package responses

import (
	"time"
)

type Corporateinfo struct {
	Branchid    int64
	Corpaddress string
	Corpid      int64
	Corpname    string
	CreatedAt   time.Time
	Email       string
	Id          int
	Logo        string
	Phone       string
	UpdatedAt   time.Time
}

type Accounts struct {
	AccountName string
	AccountNo   int
	Branchid    *Corporateinfo
	Corpid      *Corporateinfo
	DateTime    time.Time
	Id          int
	Mobile      string
}

type AccountDTO struct {
	StatusCode int
	Account    *Accounts
	StatusDesc string
}
