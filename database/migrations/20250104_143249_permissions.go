package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Permissions_20250104_143249 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Permissions_20250104_143249{}
	m.Created = "20250104_143249"

	migration.Register("Permissions_20250104_143249", m)
}

// Run the migrations
func (m *Permissions_20250104_143249) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE permissions(`permission_id` int(11) NOT NULL AUTO_INCREMENT,`permission` varchar(100) NOT NULL,`permission_code` varchar(10) NOT NULL,`permission_description` varchar(500) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT NULL,PRIMARY KEY (`permission_id`))")
}

// Reverse the migrations
func (m *Permissions_20250104_143249) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `permissions`")
}
