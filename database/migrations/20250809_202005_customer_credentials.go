package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CustomerCredentials_20250809_202005 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CustomerCredentials_20250809_202005{}
	m.Created = "20250809_202005"

	migration.Register("CustomerCredentials_20250809_202005", m)
}

// Run the migrations
func (m *CustomerCredentials_20250809_202005) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE customer_credentials(`customer_credential_id` int(11) NOT NULL AUTO_INCREMENT,`customer_id` int(11) NOT NULL,`username` varchar(255) DEFAULT NULL,`password` varchar(255) DEFAULT NULL,`pin` varchar(10) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,`active` int(11) DEFAULT 1,PRIMARY KEY (`customer_credential_id`), FOREIGN KEY (customer_id) REFERENCES customers(customer_id) ON UPDATE CASCADE ON DELETE CASCADE)")
}

// Reverse the migrations
func (m *CustomerCredentials_20250809_202005) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `customer_credentials`")
}
