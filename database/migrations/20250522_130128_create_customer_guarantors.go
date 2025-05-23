package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateCustomerGuarantors_20250522_130128 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCustomerGuarantors_20250522_130128{}
	m.Created = "20250522_130128"

	migration.Register("CreateCustomerGuarantors_20250522_130128", m)
}

// Run the migrations
func (m *CreateCustomerGuarantors_20250522_130128) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE customer_guarantors(`customer_guarantor_id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(120) NOT NULL,`contact` varchar(50) NOT NULL,`customer_id` int(11) NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`customer_guarantor_id`), FOREIGN KEY (customer_id) REFERENCES customers(customer_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *CreateCustomerGuarantors_20250522_130128) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `create_customer_guarantors`")
}
