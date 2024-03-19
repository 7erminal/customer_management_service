package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CustomerCategories_20240317_035045 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CustomerCategories_20240317_035045{}
	m.Created = "20240317_035045"

	migration.Register("CustomerCategories_20240317_035045", m)
}

// Run the migrations
func (m *CustomerCategories_20240317_035045) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE customer_categories(`customer_category_id` int(11) NOT NULL AUTO_INCREMENT,`category` varchar(100) NOT NULL,`description` varchar(255) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`customer_category_id`))")
}

// Reverse the migrations
func (m *CustomerCategories_20240317_035045) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `customer_categories`")
}
