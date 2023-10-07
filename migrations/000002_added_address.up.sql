START TRANSACTION;

CREATE TABLE addresses (
    id bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    created_at datetime(3) DEFAULT NULL,
    updated_at datetime(3) DEFAULT NULL,
    deleted_at datetime(3) DEFAULT NULL,
    street VARCHAR(255) NULL,
    city VARCHAR(255) NULL,
    `state` VARCHAR(255) NULL,
    zip VARCHAR(255) NULL,
    PRIMARY KEY (id),
    INDEX idx_address_deleted_at (deleted_at)
) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = latin1;

ALTER TABLE
    users
ADD
    COLUMN address_id BIGINT(20) UNSIGNED NULL;

ALTER TABLE
    users
ADD
    CONSTRAINT fk_users_address FOREIGN KEY (address_id) REFERENCES addresses (id);

COMMIT;
