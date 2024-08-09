package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Customers_20240317_023056 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Customers_20240317_023056{}
	m.Created = "20240317_023056"

	migration.Register("Customers_20240317_023056", m)
}

// Run the migrations
func (m *Customers_20240317_023056) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE customers(`customer_id` int(11) NOT NULL AUTO_INCREMENT,`user_id` int(11),`shop_id` int(11) DEFAULT NULL,`customer_category_id` int(11) DEFAULT NULL,`nickname` varchar(100) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT NULL,PRIMARY KEY (`customer_id`), FOREIGN KEY (user_id) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE CASCADE, FOREIGN KEY (shop_id) REFERENCES shops(shop_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *Customers_20240317_023056) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `customers`")
}
