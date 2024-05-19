package responses

type CustomersDTO struct {
	StatusCode int
	Customers  *[]interface{}
	StatusDesc string
}
