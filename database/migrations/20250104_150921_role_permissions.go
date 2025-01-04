package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type RolePermissions_20250104_150921 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RolePermissions_20250104_150921{}
	m.Created = "20250104_150921"

	migration.Register("RolePermissions_20250104_150921", m)
}

// Run the migrations
func (m *RolePermissions_20250104_150921) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE role_permissions(`role_permission_id` int(11) NOT NULL AUTO_INCREMENT,`role_id` int(11) NOT NULL,`permission_id` int(11) NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`role_permission_id`), FOREIGN KEY (role_id) REFERENCES roles(role_id) ON UPDATE CASCADE ON DELETE CASCADE, FOREIGN KEY (permission_id) REFERENCES permissions(permission_id) ON UPDATE CASCADE ON DELETE CASCADE)")
}

// Reverse the migrations
func (m *RolePermissions_20250104_150921) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `role_permissions`")
}
