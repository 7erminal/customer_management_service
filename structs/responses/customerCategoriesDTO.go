package responses

type CustomerCategoriesDTO struct {
	StatusCode         int
	CustomerCategories *[]interface{}
	StatusDesc         string
}
