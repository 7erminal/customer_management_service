package models

type UpdateUserRequestDTO struct {
	FullName      string `orm:"size(255)"`
	Username      string `orm:"size(255); omitempty; null"`
	PhoneNumber   string `orm:"size(255); omitempty; null"`
	Gender        string `orm:"size(10); omitempty; null"`
	Dob           string `orm:"size(50); omitempty; null"`
	Address       string `orm:"size(255); omitempty; null"`
	MaritalStatus string `orm:"size(255); omitempty; null"`
}
