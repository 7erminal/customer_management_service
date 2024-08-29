package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToUsersTable_20240829_190402 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToUsersTable_20240829_190402{}
	m.Created = "20240829_190402"

	migration.Register("AddColumnToUsersTable_20240829_190402", m)
}

// Run the migrations
func (m *AddColumnToUsersTable_20240829_190402) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE users ADD COLUMN image_path varChar(200) default null after username")
}

// Reverse the migrations
func (m *AddColumnToUsersTable_20240829_190402) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
