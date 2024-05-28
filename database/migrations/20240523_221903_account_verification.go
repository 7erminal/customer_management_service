package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AccountVerification_20240523_221903 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AccountVerification_20240523_221903{}
	m.Created = "20240523_221903"

	migration.Register("AccountVerification_20240523_221903", m)
}

// Run the migrations
func (m *AccountVerification_20240523_221903) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE account_verification(`account_verification_id` int(11) NOT NULL AUTO_INCREMENT,`account_id` int(11) DEFAULT NULL,`response` longtext  NOT NULL,`id_number` varchar(128) NOT NULL,`status` varchar(128) NOT NULL,`image` longtext  NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`account_verification_id`), FOREIGN KEY (account_id) REFERENCES accounts(account_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *AccountVerification_20240523_221903) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `account_verification`")
}
