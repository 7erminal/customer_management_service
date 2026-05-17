-- Customer Management Service migration generated from persistent models
-- in this project (models package), excluding DTO-only structs.

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE IF NOT EXISTS currencies (
  currency_id BIGINT NOT NULL AUTO_INCREMENT,
  symbol VARCHAR(20) NOT NULL DEFAULT '',
  currency VARCHAR(50) NOT NULL DEFAULT '',
  active INT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NULL,
  modified_by INT NULL,
  PRIMARY KEY (currency_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS countries (
  country_id BIGINT NOT NULL AUTO_INCREMENT,
  country VARCHAR(255) NOT NULL DEFAULT '',
  description VARCHAR(500) NOT NULL DEFAULT '',
  country_code VARCHAR(20) NOT NULL DEFAULT '',
  default_currency BIGINT NOT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL DEFAULT 0,
  modified_by INT NOT NULL DEFAULT 0,
  PRIMARY KEY (country_id),
  KEY idx_countries_default_currency (default_currency),
  CONSTRAINT fk_countries_default_currency FOREIGN KEY (default_currency) REFERENCES currencies(currency_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS actions (
  action_id BIGINT NOT NULL AUTO_INCREMENT,
  action VARCHAR(50) NOT NULL DEFAULT '',
  description VARCHAR(255) NOT NULL DEFAULT '',
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL DEFAULT 0,
  modified_by INT NOT NULL DEFAULT 0,
  active INT NOT NULL DEFAULT 0,
  PRIMARY KEY (action_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS roles (
  role_id BIGINT NOT NULL AUTO_INCREMENT,
  role VARCHAR(100) NOT NULL DEFAULT '',
  description VARCHAR(500) NOT NULL DEFAULT '',
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL DEFAULT 0,
  modified_by INT NOT NULL DEFAULT 0,
  active INT NOT NULL DEFAULT 0,
  PRIMARY KEY (role_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS permissions (
  permission_id BIGINT NOT NULL AUTO_INCREMENT,
  permission VARCHAR(100) NOT NULL DEFAULT '',
  permission_code VARCHAR(10) NOT NULL DEFAULT '',
  permission_description VARCHAR(500) NOT NULL DEFAULT '',
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL DEFAULT 0,
  modified_by INT NOT NULL DEFAULT 0,
  active INT NOT NULL DEFAULT 0,
  PRIMARY KEY (permission_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS shops (
  shop_id BIGINT NOT NULL AUTO_INCREMENT,
  shop_name VARCHAR(255) NOT NULL DEFAULT '',
  shop_description VARCHAR(255) NOT NULL DEFAULT '',
  shop_assistant_name VARCHAR(100) NOT NULL DEFAULT '',
  shop_assistant_number VARCHAR(100) NOT NULL DEFAULT '',
  phone_number VARCHAR(255) NOT NULL DEFAULT '',
  email VARCHAR(255) NOT NULL DEFAULT '',
  image VARCHAR(100) NOT NULL DEFAULT '',
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL DEFAULT 0,
  modified_by INT NOT NULL DEFAULT 0,
  active INT NOT NULL DEFAULT 0,
  PRIMARY KEY (shop_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS identification_types (
  identification_type_id BIGINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL DEFAULT '',
  code VARCHAR(100) NOT NULL DEFAULT '',
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL DEFAULT 0,
  modified_by INT NOT NULL DEFAULT 0,
  active INT NOT NULL DEFAULT 0,
  PRIMARY KEY (identification_type_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS customer_categories (
  customer_category_id BIGINT NOT NULL AUTO_INCREMENT,
  category VARCHAR(100) NOT NULL DEFAULT '',
  description VARCHAR(255) NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL DEFAULT 0,
  modified_by INT NOT NULL DEFAULT 0,
  active INT NOT NULL DEFAULT 0,
  PRIMARY KEY (customer_category_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS branches (
  branch_id BIGINT NOT NULL AUTO_INCREMENT,
  branch VARCHAR(80) NOT NULL DEFAULT '',
  country_id BIGINT NULL,
  location VARCHAR(255) NOT NULL DEFAULT '',
  phone_number VARCHAR(255) NOT NULL DEFAULT '',
  active INT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NULL,
  modified_by INT NULL,
  PRIMARY KEY (branch_id),
  KEY idx_branches_country_id (country_id),
  CONSTRAINT fk_branches_country FOREIGN KEY (country_id) REFERENCES countries(country_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS user_extra_details (
  user_details_id BIGINT NOT NULL AUTO_INCREMENT,
  branch BIGINT NULL,
  shop_id BIGINT NULL,
  nickname VARCHAR(100) NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL DEFAULT 0,
  modified_by INT NOT NULL DEFAULT 0,
  active INT NOT NULL DEFAULT 0,
  PRIMARY KEY (user_details_id),
  KEY idx_user_extra_details_branch (branch),
  KEY idx_user_extra_details_shop_id (shop_id),
  CONSTRAINT fk_user_extra_details_branch FOREIGN KEY (branch) REFERENCES branches(branch_id),
  CONSTRAINT fk_user_extra_details_shop FOREIGN KEY (shop_id) REFERENCES shops(shop_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS user_tokens (
  user_token_id BIGINT NOT NULL AUTO_INCREMENT,
  token VARCHAR(255) NOT NULL DEFAULT '',
  nonce VARCHAR(255) NOT NULL DEFAULT '',
  expiry_date DATETIME NOT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL DEFAULT 0,
  modified_by INT NOT NULL DEFAULT 0,
  active INT NOT NULL DEFAULT 0,
  PRIMARY KEY (user_token_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS users (
  user_id BIGINT NOT NULL AUTO_INCREMENT,
  user_details_id BIGINT NULL,
  image_path VARCHAR(200) NULL,
  user_type INT NULL,
  full_name VARCHAR(255) NOT NULL DEFAULT '',
  username VARCHAR(40) NULL,
  password VARCHAR(255) NOT NULL DEFAULT '',
  email VARCHAR(255) NULL,
  phone_number VARCHAR(255) NULL,
  gender VARCHAR(10) NOT NULL DEFAULT '',
  dob DATETIME NOT NULL,
  address VARCHAR(255) NULL,
  id_type VARCHAR(5) NULL,
  id_number VARCHAR(100) NULL,
  marital_status VARCHAR(20) NULL,
  active INT NULL,
  role BIGINT NULL,
  is_verified BOOLEAN NULL,
  date_created DATETIME NULL,
  date_modified DATETIME NULL,
  created_by INT NULL,
  modified_by INT NULL,
  PRIMARY KEY (user_id),
  KEY idx_users_user_details_id (user_details_id),
  KEY idx_users_role (role),
  CONSTRAINT fk_users_user_details FOREIGN KEY (user_details_id) REFERENCES user_extra_details(user_details_id),
  CONSTRAINT fk_users_role FOREIGN KEY (role) REFERENCES roles(role_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS user_invites (
  user_invite_id BIGINT NOT NULL AUTO_INCREMENT,
  invited_by BIGINT NOT NULL,
  invitation_token BIGINT NOT NULL,
  email VARCHAR(255) NOT NULL DEFAULT '',
  role BIGINT NOT NULL,
  status VARCHAR(255) NOT NULL DEFAULT '',
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL DEFAULT 0,
  modified_by INT NOT NULL DEFAULT 0,
  active INT NOT NULL DEFAULT 0,
  PRIMARY KEY (user_invite_id),
  KEY idx_user_invites_invited_by (invited_by),
  KEY idx_user_invites_invitation_token (invitation_token),
  KEY idx_user_invites_role (role),
  CONSTRAINT fk_user_invites_invited_by FOREIGN KEY (invited_by) REFERENCES users(user_id),
  CONSTRAINT fk_user_invites_invitation_token FOREIGN KEY (invitation_token) REFERENCES user_tokens(user_token_id),
  CONSTRAINT fk_user_invites_role FOREIGN KEY (role) REFERENCES roles(role_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS customers (
  customer_id BIGINT NOT NULL AUTO_INCREMENT,
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
  dob DATETIME NOT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL DEFAULT 0,
  modified_by INT NOT NULL DEFAULT 0,
  active INT NOT NULL DEFAULT 0,
  last_txn_date DATETIME NOT NULL,
  PRIMARY KEY (customer_id),
  KEY idx_customers_identification_type_id (identification_type_id),
  KEY idx_customers_branch (branch),
  KEY idx_customers_shop_id (shop_id),
  KEY idx_customers_customer_category_id (customer_category_id),
  CONSTRAINT fk_customers_identification_type FOREIGN KEY (identification_type_id) REFERENCES identification_types(identification_type_id),
  CONSTRAINT fk_customers_branch FOREIGN KEY (branch) REFERENCES branches(branch_id),
  CONSTRAINT fk_customers_shop FOREIGN KEY (shop_id) REFERENCES shops(shop_id),
  CONSTRAINT fk_customers_category FOREIGN KEY (customer_category_id) REFERENCES customer_categories(customer_category_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS role_permissions (
  role_permission_id BIGINT NOT NULL AUTO_INCREMENT,
  role_id BIGINT NOT NULL,
  permission_id BIGINT NOT NULL,
  action_id BIGINT NOT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL DEFAULT 0,
  modified_by INT NOT NULL DEFAULT 0,
  active INT NOT NULL DEFAULT 0,
  PRIMARY KEY (role_permission_id),
  KEY idx_role_permissions_role_id (role_id),
  KEY idx_role_permissions_permission_id (permission_id),
  KEY idx_role_permissions_action_id (action_id),
  CONSTRAINT fk_role_permissions_role FOREIGN KEY (role_id) REFERENCES roles(role_id),
  CONSTRAINT fk_role_permissions_permission FOREIGN KEY (permission_id) REFERENCES permissions(permission_id),
  CONSTRAINT fk_role_permissions_action FOREIGN KEY (action_id) REFERENCES actions(action_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS customer_emergency_contacts (
  customer_emergency_contact_id BIGINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(120) NOT NULL DEFAULT '',
  contact VARCHAR(50) NOT NULL DEFAULT '',
  customer_id BIGINT NOT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL DEFAULT 0,
  modified_by INT NOT NULL DEFAULT 0,
  PRIMARY KEY (customer_emergency_contact_id),
  KEY idx_customer_emergency_contacts_customer_id (customer_id),
  CONSTRAINT fk_customer_emergency_contacts_customer FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS customer_guarantors (
  customer_guarantor_id BIGINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(120) NOT NULL DEFAULT '',
  contact VARCHAR(50) NOT NULL DEFAULT '',
  customer_id BIGINT NOT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL DEFAULT 0,
  modified_by INT NOT NULL DEFAULT 0,
  PRIMARY KEY (customer_guarantor_id),
  KEY idx_customer_guarantors_customer_id (customer_id),
  CONSTRAINT fk_customer_guarantors_customer FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
