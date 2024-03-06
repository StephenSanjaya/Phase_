-- Creating table + data
CREATE TABLE Employees (
    employee_id INT PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    position VARCHAR(50)
);

CREATE TABLE MenuItems (
    item_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50),
    description VARCHAR(255),
    price INT,
    category VARCHAR(50)
);

CREATE TABLE Orders (
    order_id INT PRIMARY KEY AUTO_INCREMENT,
    table_number INT,
    employee_id INT,
    order_date DATE,
    status VARCHAR(50),
    FOREIGN KEY (employee_id) REFERENCES Employees(employee_id)
);

CREATE TABLE OrderItems (
    order_item_id INT PRIMARY KEY AUTO_INCREMENT,
    order_id INT,
    item_id INT,
    quantity INT,
    subtotal INT,
    FOREIGN KEY (order_id) REFERENCES Orders(order_id),
    FOREIGN KEY (item_id) REFERENCES MenuItems(item_id)
);

CREATE TABLE Payments (
    payment_id INT PRIMARY KEY AUTO_INCREMENT,
    order_id INT,
    payment_date DATE,
    payment_method VARCHAR(50),
    total_amount INT,
    FOREIGN KEY (order_id) REFERENCES Orders(order_id)
);

INSERT INTO Employees (first_name, last_name, position)
VALUES ("John", "Doe", "Waiter");

INSERT INTO MenuItems (name, description, price, category)
VALUES ("Steak", "Grilled sirloin steak", 25.99, "Main course");

INSERT INTO Orders (table_number, employee_id, order_date, status)
VALUES (5, 1, "2023-08-04", "Pending");

INSERT INTO OrderItems (order_id, item_id, quantity, subtotal)
VALUES (1, 1, 2, 51.98);

INSERT INTO Payments (order_id, payment_date, payment_method, total_amount)
VALUES (1, "2023-08-04", "Credit Card", 51.98);



-- A. Retrieve all orders with their applied discounts
SELECT oi.order_id, oi.item_id, oi.subtotal, 
CASE
    WHEN oi.subtotal > 50.00 THEN oi.subtotal * 0.2
    ELSE oi.subtotal
END AS Discount
FROM OrderItems oi
JOIN Orders o ON oi.order_id = o.order_id

-- B. Calculate the total revenue (including discounts) for a specific day
SELECT SUM(CASE WHEN subtotal > 50.00 THEN subtotal - (subtotal * 0.2) ELSE subtotal END) AS "Total Revenue"
FROM OrderItems