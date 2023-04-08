CREATE TABLE orders_item (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(1024),
    price int,
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);