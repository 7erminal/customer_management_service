package models

type CustomerCategoriesRequestDTO struct {
	Category    string `orm:"size(100)"`
	Description string `orm:"size(255); null"`
	CreatedBy   int
}

type CustomerCategoriesResponseDTO struct {
	StatusCode int
	Category   *Customer_categories
	StatusDesc string `orm:"size(255)"`
}
