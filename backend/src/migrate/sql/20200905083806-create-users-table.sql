
-- +migrate Up
CREATE TABLE users (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `email` VARCHAR(200) NOT NULL,
  `first_name` VARCHAR(100) NOT NULL,
  `last_name` VARCHAR(100) NOT NULL,
  `display_name` VARCHAR(50) NOT NULL,
  `last_signin_at` DATETIME,
  `created_at` DATETIME,
  `updated_at` DATETIME,
  `deleted_at` DATETIME
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

CREATE UNIQUE INDEX index_of_unique_users_on_email ON users(email);

-- +migrate Down
DROP TABLE users;
