package responses

import "time"

type Currencies struct {
	CurrencyId   int64
	Symbol       string
	Currency     string
	Active       int
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    int
	ModifiedBy   int
}

type Countries struct {
	CountryId       int64
	Country         string
	Description     string
	CountryCode     string
	DefaultCurrency *Currencies
	DateCreated     time.Time
	DateModified    time.Time
	CreatedBy       int
	ModifiedBy      int
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
