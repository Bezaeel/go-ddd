-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    shipment_number INTEGER,
    cargo_id INTEGER,
    is_shipped BOOLEAN,
    created_at TIMESTAMP
);

CREATE TABLE order_line_items (
    id SERIAL PRIMARY KEY,
    product_id INTEGER,
    seller_id INTEGER,
    order_id INTEGER,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_line_items;
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
