package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Accounts_20240523_202727 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Accounts_20240523_202727{}
	m.Created = "20240523_202727"

	migration.Register("Accounts_20240523_202727", m)
}

// Run the migrations
func (m *Accounts_20240523_202727) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE accounts(`account_id` int(11) NOT NULL AUTO_INCREMENT,`user_id` int(11) DEFAULT NULL,`account_number` varchar(255) NOT NULL,`balance` float DEFAULT 0,`balance_before` float DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`account_id`), FOREIGN KEY (user_id) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *Accounts_20240523_202727) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `accounts`")
}
