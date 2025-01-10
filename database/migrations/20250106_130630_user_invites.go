package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type UserInvites_20250106_130630 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserInvites_20250106_130630{}
	m.Created = "20250106_130630"

	migration.Register("UserInvites_20250106_130630", m)
}

// Run the migrations
func (m *UserInvites_20250106_130630) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE user_invites(`user_invite_id` int(11) NOT NULL AUTO_INCREMENT,`invitation_token` int(11) NOT NULL,`invited_by` int(11) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`user_invite_id`), FOREIGN KEY (invited_by) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE CASCADE, FOREIGN KEY (invitation_token) REFERENCES user_tokens(user_token_id) ON UPDATE CASCADE ON DELETE CASCADE)")
}

// Reverse the migrations
func (m *UserInvites_20250106_130630) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user_invites`")
}
