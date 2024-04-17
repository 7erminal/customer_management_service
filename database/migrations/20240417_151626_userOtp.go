package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type UserOtp_20240417_151626 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserOtp_20240417_151626{}
	m.Created = "20240417_151626"

	migration.Register("UserOtp_20240417_151626", m)
}

// Run the migrations
func (m *UserOtp_20240417_151626) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE user_otps(`user_otp_id` int(11) NOT NULL AUTO_INCREMENT,`code` varchar(128) NOT NULL,`user_id` int(11) NOT NULL,`status` int(11) DEFAULT 2,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_generated` datetime DEFAULT CURRENT_TIMESTAMP,`expiry_date` datetime DEFAULT NULL,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`user_otp_id`), FOREIGN KEY (user_id) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *UserOtp_20240417_151626) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `userOtp`")
}
