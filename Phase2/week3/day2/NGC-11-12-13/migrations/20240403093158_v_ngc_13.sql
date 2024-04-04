-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stores (
    store_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    longitude VARCHAR(100) NOT NULL,
    latitude VARCHAR(100) NOT NULL,
    rating DECIMAL(10,1) NOT NULL
);

ALTER TABLE products
ADD store_id int,
ADD FOREIGN KEY (store_id) REFERENCES stores(store_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE products DROP FOREIGN KEY store_id;

ALTER TABLE products DROP store_id;

DROP TABLE stores;
-- +goose StatementEnd
