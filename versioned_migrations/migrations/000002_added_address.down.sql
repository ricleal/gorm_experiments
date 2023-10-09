START TRANSACTION;

ALTER TABLE
    users DROP FOREIGN KEY fk_users_address;

ALTER TABLE
    users DROP COLUMN address_id;

DROP TABLE IF EXISTS `addresses`;

COMMIT;
