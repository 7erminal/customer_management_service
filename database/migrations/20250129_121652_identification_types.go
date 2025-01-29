package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type IdentificationTypes_20250129_121652 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &IdentificationTypes_20250129_121652{}
	m.Created = "20250129_121652"

	migration.Register("IdentificationTypes_20250129_121652", m)
}

// Run the migrations
func (m *IdentificationTypes_20250129_121652) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE identification_types(`identification_type_id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(100) NOT NULL,`code` varchar(100) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`identification_type_id`))")
}

// Reverse the migrations
func (m *IdentificationTypes_20250129_121652) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `identification_types`")
}
