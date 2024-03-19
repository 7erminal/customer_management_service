package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Shops_20240317_022536 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Shops_20240317_022536{}
	m.Created = "20240317_022536"

	migration.Register("Shops_20240317_022536", m)
}

// Run the migrations
func (m *Shops_20240317_022536) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE shops(`shop_id` int(11) NOT NULL AUTO_INCREMENT,`shop_name` varchar(100),`shop_description` varchar(255) DEFAULT NULL,`shop_assistant_name` varchar(100) DEFAULT NULL,`shop_assistant_number` varchar(100) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT NULL,PRIMARY KEY (`shop_id`))")
}

// Reverse the migrations
func (m *Shops_20240317_022536) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `shops`")
}
