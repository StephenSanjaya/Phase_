CREATE TABLE Customers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    address VARCHAR(200) NOT NULL
);

CREATE TABLE Products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price NUMERIC(10,2) NOT NULL,
    inventory_level INTEGER NOT NULL
);

INSERT INTO products (name, description, price, inventory_level) 
VALUES ('Coca Cola', 'Minuman soda', 10.25, 1), ('Teh Pucuk', 'Minuman teh', 6.25, 1), ('Teh Botol', 'Minuman teh', 7.25, 1);

CREATE TABLE Orders (
    id SERIAL PRIMARY KEY,
    order_date DATE NOT NULL,
    customer_id INTEGER NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    payment_confirmation VARCHAR(100) NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES Customers (customer_id)
);

CREATE TABLE Order_Items (
    order_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES Orders (order_id),
    FOREIGN KEY (product_id) REFERENCES Products (product_id)
);

CREATE TABLE Payments (
    payment_id SERIAL PRIMARY KEY,
    order_id INTEGER UNIQUE NOT NULL,
    amount NUMERIC(10,2) NOT NULL,
    payment_date DATE NOT NULL,
    FOREIGN KEY (order_id) REFERENCES Orders (order_id)
);