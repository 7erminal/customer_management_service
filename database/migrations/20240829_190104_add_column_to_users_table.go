package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToUsersTable_20240829_190104 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToUsersTable_20240829_190104{}
	m.Created = "20240829_190104"

	migration.Register("AddColumnToUsersTable_20240829_190104", m)
}

// Run the migrations
func (m *AddColumnToUsersTable_20240829_190104) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE users ADD COLUMN ImagePath varChar(200) default null after username")
}

// Reverse the migrations
func (m *AddColumnToUsersTable_20240829_190104) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
