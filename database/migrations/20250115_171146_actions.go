package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Actions_20250115_171146 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Actions_20250115_171146{}
	m.Created = "20250115_171146"

	migration.Register("Actions_20250115_171146", m)
}

// Run the migrations
func (m *Actions_20250115_171146) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE actions(`action_id` int(11) NOT NULL AUTO_INCREMENT,`action` varchar(50) NOT NULL,`description` varchar(255) NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`action_id`))")
}

// Reverse the migrations
func (m *Actions_20250115_171146) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `actions`")
}
