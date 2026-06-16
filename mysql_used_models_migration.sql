-- Customer Management Service migration generated from persistent models
-- in this project (models package), excluding DTO-only structs.

CREATE TABLE IF NOT EXISTS actions (
  action_id BIGSERIAL PRIMARY KEY,
  action VARCHAR(50) NOT NULL DEFAULT '',
  description VARCHAR(255) NOT NULL DEFAULT '',
  date_created TIMESTAMP NOT NULL,
  date_modified TIMESTAMP NOT NULL,
  created_by INTEGER NOT NULL DEFAULT 0,
  modified_by INTEGER NOT NULL DEFAULT 0,
  active INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS roles (
  role_id BIGSERIAL PRIMARY KEY,
  role VARCHAR(100) NOT NULL DEFAULT '',
  description VARCHAR(500) NOT NULL DEFAULT '',
  date_created TIMESTAMP NOT NULL,
  date_modified TIMESTAMP NOT NULL,
  created_by INTEGER NOT NULL DEFAULT 0,
  modified_by INTEGER NOT NULL DEFAULT 0,
  active INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS permissions (
  permission_id BIGSERIAL PRIMARY KEY,
  permission VARCHAR(100) NOT NULL DEFAULT '',
  permission_code VARCHAR(10) NOT NULL DEFAULT '',
  permission_description VARCHAR(500) NOT NULL DEFAULT '',
  date_created TIMESTAMP NOT NULL,
  date_modified TIMESTAMP NOT NULL,
  created_by INTEGER NOT NULL DEFAULT 0,
  modified_by INTEGER NOT NULL DEFAULT 0,
  active INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS shops (
  shop_id BIGSERIAL PRIMARY KEY,
  shop_name VARCHAR(255) NOT NULL DEFAULT '',
  shop_description VARCHAR(255) NOT NULL DEFAULT '',
  shop_assistant_name VARCHAR(100) NOT NULL DEFAULT '',
  shop_assistant_number VARCHAR(100) NOT NULL DEFAULT '',
  phone_number VARCHAR(255) NOT NULL DEFAULT '',
  email VARCHAR(255) NOT NULL DEFAULT '',
  image VARCHAR(100) NOT NULL DEFAULT '',
  date_created TIMESTAMP NOT NULL,
  date_modified TIMESTAMP NOT NULL,
  created_by INTEGER NOT NULL DEFAULT 0,
  modified_by INTEGER NOT NULL DEFAULT 0,
  active INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS identification_types (
  identification_type_id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL DEFAULT '',
  code VARCHAR(100) NOT NULL DEFAULT '',
  date_created TIMESTAMP NOT NULL,
  date_modified TIMESTAMP NOT NULL,
  created_by INTEGER NOT NULL DEFAULT 0,
  modified_by INTEGER NOT NULL DEFAULT 0,
  active INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS customer_categories (
  customer_category_id BIGSERIAL PRIMARY KEY,
  category VARCHAR(100) NOT NULL DEFAULT '',
  description VARCHAR(255) NULL,
  date_created TIMESTAMP NOT NULL,
  date_modified TIMESTAMP NOT NULL,
  created_by INTEGER NOT NULL DEFAULT 0,
  modified_by INTEGER NOT NULL DEFAULT 0,
  active INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS branches (
  branch_id BIGSERIAL PRIMARY KEY,
  branch VARCHAR(80) NOT NULL DEFAULT '',
  country_id BIGINT NULL,
  location VARCHAR(255) NOT NULL DEFAULT '',
  phone_number VARCHAR(255) NOT NULL DEFAULT '',
  active INTEGER NULL,
  date_created TIMESTAMP NOT NULL,
  date_modified TIMESTAMP NOT NULL,
  created_by INTEGER NULL,
  modified_by INTEGER NULL
);
CREATE INDEX IF NOT EXISTS idx_branches_country_id ON branches(country_id);

CREATE TABLE IF NOT EXISTS user_extra_details (
  user_details_id BIGSERIAL PRIMARY KEY,
  branch BIGINT NULL,
  shop_id BIGINT NULL,
  nickname VARCHAR(100) NULL,
  date_created TIMESTAMP NOT NULL,
  date_modified TIMESTAMP NOT NULL,
  created_by INTEGER NOT NULL DEFAULT 0,
  modified_by INTEGER NOT NULL DEFAULT 0,
  active INTEGER NOT NULL DEFAULT 0,
  CONSTRAINT fk_user_extra_details_branch FOREIGN KEY (branch) REFERENCES branches(branch_id),
  CONSTRAINT fk_user_extra_details_shop FOREIGN KEY (shop_id) REFERENCES shops(shop_id)
);
CREATE INDEX IF NOT EXISTS idx_user_extra_details_branch ON user_extra_details(branch);
CREATE INDEX IF NOT EXISTS idx_user_extra_details_shop_id ON user_extra_details(shop_id);

CREATE TABLE IF NOT EXISTS user_tokens (
  user_token_id BIGSERIAL PRIMARY KEY,
  token VARCHAR(255) NOT NULL DEFAULT '',
  nonce VARCHAR(255) NOT NULL DEFAULT '',
  expiry_date TIMESTAMP NOT NULL,
  date_created TIMESTAMP NOT NULL,
  date_modified TIMESTAMP NOT NULL,
  created_by INTEGER NOT NULL DEFAULT 0,
  modified_by INTEGER NOT NULL DEFAULT 0,
  active INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS users (
  user_id BIGSERIAL PRIMARY KEY,
  user_details_id BIGINT NULL,
  image_path VARCHAR(200) NULL,
  user_type INTEGER NULL,
  full_name VARCHAR(255) NOT NULL DEFAULT '',
  username VARCHAR(40) NULL,
  password VARCHAR(255) NOT NULL DEFAULT '',
  email VARCHAR(255) NULL,
  phone_number VARCHAR(255) NULL,
  gender VARCHAR(10) NOT NULL DEFAULT '',
  dob TIMESTAMP NOT NULL,
  address VARCHAR(255) NULL,
  id_type VARCHAR(5) NULL,
  id_number VARCHAR(100) NULL,
  marital_status VARCHAR(20) NULL,
  active INTEGER NULL,
  role BIGINT NULL,
  is_verified BOOLEAN NULL,
  date_created TIMESTAMP NULL,
  date_modified TIMESTAMP NULL,
  created_by INTEGER NULL,
  modified_by INTEGER NULL,
  CONSTRAINT fk_users_user_details FOREIGN KEY (user_details_id) REFERENCES user_extra_details(user_details_id),
  CONSTRAINT fk_users_role FOREIGN KEY (role) REFERENCES roles(role_id)
);
CREATE INDEX IF NOT EXISTS idx_users_user_details_id ON users(user_details_id);
CREATE INDEX IF NOT EXISTS idx_users_role ON users(role);

CREATE TABLE IF NOT EXISTS user_invites (
  user_invite_id BIGSERIAL PRIMARY KEY,
  invited_by BIGINT NOT NULL,
  invitation_token BIGINT NOT NULL,
  email VARCHAR(255) NOT NULL DEFAULT '',
  role BIGINT NOT NULL,
  status VARCHAR(255) NOT NULL DEFAULT '',
  date_created TIMESTAMP NOT NULL,
  date_modified TIMESTAMP NOT NULL,
  created_by INTEGER NOT NULL DEFAULT 0,
  modified_by INTEGER NOT NULL DEFAULT 0,
  active INTEGER NOT NULL DEFAULT 0,
  CONSTRAINT fk_user_invites_invited_by FOREIGN KEY (invited_by) REFERENCES users(user_id),
  CONSTRAINT fk_user_invites_invitation_token FOREIGN KEY (invitation_token) REFERENCES user_tokens(user_token_id),
  CONSTRAINT fk_user_invites_role FOREIGN KEY (role) REFERENCES roles(role_id)
);
CREATE INDEX IF NOT EXISTS idx_user_invites_invited_by ON user_invites(invited_by);
CREATE INDEX IF NOT EXISTS idx_user_invites_invitation_token ON user_invites(invitation_token);
CREATE INDEX IF NOT EXISTS idx_user_invites_role ON user_invites(role);

CREATE TABLE IF NOT EXISTS customers (
  customer_id BIGSERIAL PRIMARY KEY,
  customer_number VARCHAR(255) NOT NULL DEFAULT '',
  full_name VARCHAR(255) NOT NULL DEFAULT '',
  image_path VARCHAR(255) NOT NULL DEFAULT '',
  email VARCHAR(255) NULL,
  phone_number VARCHAR(255) NULL,
  gender VARCHAR(10) NULL,
  location VARCHAR(255) NULL,
  identification_type_id BIGINT NULL,
  identification_number VARCHAR(255) NULL,
  branch BIGINT NULL,
  shop_id BIGINT NULL,
  customer_category_id BIGINT NULL,
  nickname VARCHAR(100) NULL,
  dob TIMESTAMP NOT NULL,
  date_created TIMESTAMP NOT NULL,
  date_modified TIMESTAMP NOT NULL,
  created_by INTEGER NOT NULL DEFAULT 0,
  modified_by INTEGER NOT NULL DEFAULT 0,
  active INTEGER NOT NULL DEFAULT 0,
  last_txn_date TIMESTAMP NOT NULL,
  CONSTRAINT fk_customers_identification_type FOREIGN KEY (identification_type_id) REFERENCES identification_types(identification_type_id),
  CONSTRAINT fk_customers_branch FOREIGN KEY (branch) REFERENCES branches(branch_id),
  CONSTRAINT fk_customers_shop FOREIGN KEY (shop_id) REFERENCES shops(shop_id),
  CONSTRAINT fk_customers_category FOREIGN KEY (customer_category_id) REFERENCES customer_categories(customer_category_id)
);
CREATE INDEX IF NOT EXISTS idx_customers_identification_type_id ON customers(identification_type_id);
CREATE INDEX IF NOT EXISTS idx_customers_branch ON customers(branch);
CREATE INDEX IF NOT EXISTS idx_customers_shop_id ON customers(shop_id);
CREATE INDEX IF NOT EXISTS idx_customers_customer_category_id ON customers(customer_category_id);

CREATE TABLE IF NOT EXISTS role_permissions (
  role_permission_id BIGSERIAL PRIMARY KEY,
  role_id BIGINT NOT NULL,
  permission_id BIGINT NOT NULL,
  action_id BIGINT NOT NULL,
  date_created TIMESTAMP NOT NULL,
  date_modified TIMESTAMP NOT NULL,
  created_by INTEGER NOT NULL DEFAULT 0,
  modified_by INTEGER NOT NULL DEFAULT 0,
  active INTEGER NOT NULL DEFAULT 0,
  CONSTRAINT fk_role_permissions_role FOREIGN KEY (role_id) REFERENCES roles(role_id),
  CONSTRAINT fk_role_permissions_permission FOREIGN KEY (permission_id) REFERENCES permissions(permission_id),
  CONSTRAINT fk_role_permissions_action FOREIGN KEY (action_id) REFERENCES actions(action_id)
);
CREATE INDEX IF NOT EXISTS idx_role_permissions_role_id ON role_permissions(role_id);
CREATE INDEX IF NOT EXISTS idx_role_permissions_permission_id ON role_permissions(permission_id);
CREATE INDEX IF NOT EXISTS idx_role_permissions_action_id ON role_permissions(action_id);

CREATE TABLE IF NOT EXISTS customer_emergency_contacts (
  customer_emergency_contact_id BIGSERIAL PRIMARY KEY,
  name VARCHAR(120) NOT NULL DEFAULT '',
  contact VARCHAR(50) NOT NULL DEFAULT '',
  customer_id BIGINT NOT NULL,
  date_created TIMESTAMP NOT NULL,
  date_modified TIMESTAMP NOT NULL,
  created_by INTEGER NOT NULL DEFAULT 0,
  modified_by INTEGER NOT NULL DEFAULT 0,
  CONSTRAINT fk_customer_emergency_contacts_customer FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);
CREATE INDEX IF NOT EXISTS idx_customer_emergency_contacts_customer_id ON customer_emergency_contacts(customer_id);

CREATE TABLE IF NOT EXISTS customer_guarantors (
  customer_guarantor_id BIGSERIAL PRIMARY KEY,
  name VARCHAR(120) NOT NULL DEFAULT '',
  contact VARCHAR(50) NOT NULL DEFAULT '',
  customer_id BIGINT NOT NULL,
  date_created TIMESTAMP NOT NULL,
  date_modified TIMESTAMP NOT NULL,
  created_by INTEGER NOT NULL DEFAULT 0,
  modified_by INTEGER NOT NULL DEFAULT 0,
  CONSTRAINT fk_customer_guarantors_customer FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);
CREATE INDEX IF NOT EXISTS idx_customer_guarantors_customer_id ON customer_guarantors(customer_id);
