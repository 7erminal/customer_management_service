package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToCustomersTable_20250102_170231 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToCustomersTable_20250102_170231{}
	m.Created = "20250102_170231"

	migration.Register("AddColumnToCustomersTable_20250102_170231", m)
}

// Run the migrations
func (m *AddColumnToCustomersTable_20250102_170231) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE customers add COLUMN branch int(11) default NULL, ADD FOREIGN KEY (branch) REFERENCES branches(branch_id) ON UPDATE CASCADE ON DELETE NO ACTION")
}

// Reverse the migrations
func (m *AddColumnToCustomersTable_20250102_170231) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
