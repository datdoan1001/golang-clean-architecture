-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
   id INT NOT NULL AUTO_INCREMENT,
   name varchar(255) DEFAULT NULL COMMENT 'User name',
   age varchar(255) DEFAULT NULL COMMENT 'Age',
   created_at datetime DEFAULT NULL COMMENT 'Creation date and time',
   updated_at datetime DEFAULT NULL COMMENT 'Update date and time',
   deleted_at timestamp NULL DEFAULT NULL COMMENT 'Delete date and time',
   INDEX user_id (id),
   PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8 COMMENT='users';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
