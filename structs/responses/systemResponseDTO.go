package responses

import "time"

type Currencies struct {
	CurrencyId   string
	Symbol       string
	Currency     string
	Active       int
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    string
	ModifiedBy   string
}

type Countries struct {
	CountryId       string
	Country         string
	Description     string
	CountryCode     string
	DefaultCurrency *Currencies
	DateCreated     time.Time
	DateModified    time.Time
	CreatedBy       string
	ModifiedBy      string
}

type CountryResponseDTO struct {
	StatusCode int
	Result     *Countries
	StatusDesc string
}

type CurrencyResponseDTO struct {
	StatusCode int
	Result     *Currencies
	StatusDesc string
}
