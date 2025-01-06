package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type UserInvites_20250105_012619 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserInvites_20250105_012619{}
	m.Created = "20250105_012619"

	migration.Register("UserInvites_20250105_012619", m)
}

// Run the migrations
func (m *UserInvites_20250105_012619) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE user_invites(`user_invite_id` int(11) NOT NULL AUTO_INCREMENT,`invited_by` int(11) NOT NULL,`invitation_token` varchar(255) NOT NULL,`expiry_date` datetime NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,`active` int(11) DEFAULT NULL,PRIMARY KEY (`user_invite_id`), FOREIGN KEY (invited_by) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE CASCADE)")
}

// Reverse the migrations
func (m *UserInvites_20250105_012619) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user_invites`")
}
