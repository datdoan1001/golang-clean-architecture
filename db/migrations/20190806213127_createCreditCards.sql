
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE credit_cards (
   id INT NOT NULL AUTO_INCREMENT,
   user_id INT NOT NULL,
   number varchar(255) DEFAULT NULL COMMENT 'User name',
   created_at datetime DEFAULT NULL COMMENT 'Creation date and time',
   updated_at datetime DEFAULT NULL COMMENT 'Update date and time',
   deleted_at timestamp NULL DEFAULT NULL COMMENT 'Delete date and time',
   PRIMARY KEY(id),
   CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf COMMENT='credit card';

-- +goose Down
DROP TABLE credit_cards
-- SQL section 'Down' is executed when this migration is rolled back

