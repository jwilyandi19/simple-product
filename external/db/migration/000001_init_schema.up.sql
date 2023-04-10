CREATE TABLE orders_item (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(1024),
    price int,
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    full_name varchar(1024),
    first_order varchar(1024),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE order_histories (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    order_item_id int NOT NULL,
    descriptions varchar(1024),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE order_histories ADD FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE order_histories ADD FOREIGN KEY (order_item_id) REFERENCES orders_item(id);