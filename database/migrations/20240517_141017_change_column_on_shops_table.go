package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ChangeColumnOnShopsTable_20240517_141017 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ChangeColumnOnShopsTable_20240517_141017{}
	m.Created = "20240517_141017"

	migration.Register("ChangeColumnOnShopsTable_20240517_141017", m)
}

// Run the migrations
func (m *ChangeColumnOnShopsTable_20240517_141017) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE shops MODIFY image varchar(255)")
}

// Reverse the migrations
func (m *ChangeColumnOnShopsTable_20240517_141017) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
