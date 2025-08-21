package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToCustomersTable_20250815_094435 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToCustomersTable_20250815_094435{}
	m.Created = "20250815_094435"

	migration.Register("AddColumnToCustomersTable_20250815_094435", m)
}

// Run the migrations
func (m *AddColumnToCustomersTable_20250815_094435) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE customers ADD COLUMN customer_number varchar(255) DEFAULT NULL UNIQUE AFTER customer_id")
}

// Reverse the migrations
func (m *AddColumnToCustomersTable_20250815_094435) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
