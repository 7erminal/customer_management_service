package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Roles_20250104_141428 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Roles_20250104_141428{}
	m.Created = "20250104_141428"

	migration.Register("Roles_20250104_141428", m)
}

// Run the migrations
func (m *Roles_20250104_141428) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE roles(`role_id` int(11) NOT NULL AUTO_INCREMENT,`role` varchar(100) NOT NULL,`description` varchar(500) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`role_id`))")
}

// Reverse the migrations
func (m *Roles_20250104_141428) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `roles`")
}
