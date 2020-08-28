
-- +migrate Up
CREATE TABLE todos (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(100) NOT NULL,
  `priority`  VARCHAR(20) DEFAULT 'undefined',
  `description` TEXT,
  `created_at` DATETIME,
  `updated_at` DATETIME,
  `deleted_at` DATETIME
) ENGINE = InnoDB DEFAULT CHARSET=utf8;
-- +migrate Down
DROP TABLE todos;
