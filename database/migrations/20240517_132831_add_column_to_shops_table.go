package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToShopsTable_20240517_132831 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToShopsTable_20240517_132831{}
	m.Created = "20240517_132831"

	migration.Register("AddColumnToShopsTable_20240517_132831", m)
}

// Run the migrations
func (m *AddColumnToShopsTable_20240517_132831) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE shops add image text default null after shop_assistant_number")
}

// Reverse the migrations
func (m *AddColumnToShopsTable_20240517_132831) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
