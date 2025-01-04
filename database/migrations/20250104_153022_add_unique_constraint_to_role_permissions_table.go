package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddUniqueConstraintToRolePermissionsTable_20250104_153022 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddUniqueConstraintToRolePermissionsTable_20250104_153022{}
	m.Created = "20250104_153022"

	migration.Register("AddUniqueConstraintToRolePermissionsTable_20250104_153022", m)
}

// Run the migrations
func (m *AddUniqueConstraintToRolePermissionsTable_20250104_153022) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE `role_permissions` ADD UNIQUE `unique_index`(`role_id`, `permission_id`)")
}

// Reverse the migrations
func (m *AddUniqueConstraintToRolePermissionsTable_20250104_153022) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
