CREATE TABLE IF NOT EXISTS Employees (
    employee_id INT AUTO_INCREMENT,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    position VARCHAR(50),
    PRIMARY KEY (employee_id)
);

CREATE TABLE IF NOT EXISTS MenuItems (
    item_id INT AUTO_INCREMENT,
    name VARCHAR(50),
    description VARCHAR(255),
    price DECIMAL(10, 2),
    category VARCHAR(50),
    PRIMARY KEY (item_id)
);

CREATE TABLE IF NOT EXISTS Orders (
    order_id INT PRIMARY KEY AUTO_INCREMENT,
    table_number INT,
    employee_id INT,
    order_date DATE,
    status VARCHAR(50),
    FOREIGN KEY (employee_id) REFERENCES Employees(employee_id)
);

CREATE TABLE IF NOT EXISTS OrderItems (
    order_item_id INT PRIMARY KEY AUTO_INCREMENT,
    order_id INT,
    item_id INT,
    quantity INT,
    subtotal DECIMAL(10, 2),
    FOREIGN KEY (order_id) REFERENCES Orders(order_id),
    FOREIGN KEY (item_id) REFERENCES MenuItems(item_id)
);

CREATE TABLE IF NOT EXISTS Payments (
    payment_id INT PRIMARY KEY AUTO_INCREMENT,
    order_id INT,
    payment_date DATE,
    payment_method VARCHAR(50),
    total_amount DECIMAL(10, 2),
    FOREIGN KEY (order_id) REFERENCES Orders(order_id)
);

INSERT INTO Employees (first_name, last_name, position) 
VALUES  ("John", "Doe", "Manager"),
        ("Jane", "Smith", "Waitress"),
        ("Robert", "Brown", "Cook");

INSERT INTO MenuItems (name, description, price, category)
VALUES  ("Spaghetti carbonara", "Traditional Italian dish with eggs, cheese pancetta, and pepper.", 12.50, "Main Course"),
        ("Caesar Salad", "Fresh lettuce with Caesar dressing, croutons and parmesan.", 6.00, "Starter"),
        ("Tiramisu", "Classic Italian dessert with coffe-soaked sponge and mascarpone.", 5.50, "Dessert");

INSERT INTO Orders (table_number, employee_id, order_date, status) 
VALUES  (10, 1, "2023-08-09", "Placed"),
        (5, 2, "2023-08-09", "Completed");

INSERT INTO OrderItems (order_id, item_id, quantity, subtotal)
VALUES  (1, 1, 2, 25.00),
        (2, 2, 1, 6.00),
        (2, 3, 1, 5.50);

INSERT INTO Payments (order_id, payment_date, payment_method, total_amount)
VALUES  (1, "2023-08-09", "Credit Card", 25.00),
        (2, "2023-08-09", "Cash", 11.50);
        