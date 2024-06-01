package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type NewsletterCustomers_20240601_091752 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &NewsletterCustomers_20240601_091752{}
	m.Created = "20240601_091752"

	migration.Register("NewsletterCustomers_20240601_091752", m)
}

// Run the migrations
func (m *NewsletterCustomers_20240601_091752) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE newsletter_customers(`customer_id` int(11) NOT NULL AUTO_INCREMENT,`first_name` varchar(100) DEFAULT NULL,`last_name` varchar(100) DEFAULT NULL,`email` varchar(100) NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`customer_id`))")
}

// Reverse the migrations
func (m *NewsletterCustomers_20240601_091752) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `newsletter_customers`")
}
