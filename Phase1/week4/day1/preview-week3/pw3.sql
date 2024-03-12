-- 1NF
CREATE TABLE Bookstores (){
    order_id INT PRIMARY KEY AUTO_INCREMENT,
    customer_name VARCHAR(100),
    customer_email VARCHAR(100),
    book_title VARCHAR(100),
    book_type VARCHAR(100),
    author_name VARCHAR(100),
    author_email VARCHAR(100),
    order_date DATE,
    price DECIMAL(10, 2),
}

-- 2NF
CREATE TABLE IF NOT EXISTS Orders (
    order_id INT PRIMARY KEY AUTO_INCREMENT,
    order_date DATE,
);
CREATE TABLE IF NOT EXISTS Customers (
    customer_name VARCHAR(100),
    customer_email VARCHAR(100)
);
CREATE TABLE IF NOT EXISTS Authors (
    author_name VARCHAR(100),
    author_email VARCHAR(100)
);
CREATE TABLE IF NOT EXISTS Books (
    book_title VARCHAR(100),
    book_type VARCHAR(100)
);

-- 3NF
CREATE TABLE IF NOT EXISTS Customers (
    customer_id INT PRIMARY KEY AUTO_INCREMENT,
    customer_name VARCHAR(100),
    customer_email VARCHAR(100)
);
CREATE TABLE IF NOT EXISTS Authors (
    author_id INT PRIMARY KEY AUTO_INCREMENT,
    author_name VARCHAR(100),
    author_email VARCHAR(100)
);
CREATE TABLE IF NOT EXISTS Books (
    book_id INT PRIMARY KEY AUTO_INCREMENT,
    author_id INT,
    book_title VARCHAR(100),
    book_type VARCHAR(100),
    price DECIMAL(10, 2),
    FOREIGN KEY (author_id) REFERENCES Authors(author_id)
);
CREATE TABLE IF NOT EXISTS Orders (
    order_id INT PRIMARY KEY AUTO_INCREMENT,
    customer_id INT,
    order_date DATE,
    FOREIGN KEY (customer_id) REFERENCES Customers(customer_id)
);
CREATE TABLE IF NOT EXISTS Order_Details (
    order_detail_id INT PRIMARY KEY AUTO_INCREMENT,
    book_id INT,
    order_id INT,
    FOREIGN KEY (order_id) REFERENCES Orders(order_id),
    FOREIGN KEY (book_id) REFERENCES Books(book_id)
);

INSERT INTO Customers (customer_name, customer_email)
VALUES  ("John Doe", "john.doe@emai.com"),
        ("Alice Bob", "alice.bob@emai.com");

INSERT INTO Authors (author_name, author_email)
VALUES  ("Jane Smith", "jane.smith@emai.com"),
        ("Tom Brown", "tom.brown@emai.com");

INSERT INTO Books (author_id, book_title, book_type, price)
VALUES  (1, "Database Design", "Physical", 25.99),
        (1, "Database Design", "E-Book", 20.99),
        (2, "Web Development", "E-Book", 19.99);

INSERT INTO Orders (customer_id, order_date)
VALUES  (1, "2023-08-10"),
        (1, "2023-08-11"),
        (2, "2023-08-12");

INSERT INTO Order_Details (order_id, book_id)
VALUES  (1, 1),
        (2, 3),
        (3, 2);

-- a. List all books by 'Jane Smith'.
SELECT b.book_title
FROM Books b
JOIN Authors a ON b.author_id = a.author_id
WHERE a.author_name = "Jane Smith"

-- b. Find the total sales (in terms of price) for each book type.     
SELECT b.book_type, SUM(b.price)
FROM Order_Details od
JOIN Books b ON od.book_id = b.book_id
GROUP BY b.book_type

-- c. Identify customers who have ordered more than one book.     
SELECT c.customer_name, COUNT(o.customer_id)
FROM Orders o 
JOIN Customers c ON o.customer_id = c.customer_id
GROUP BY c.customer_name
HAVING COUNT(o.customer_id) > 1

-- d. Display the author who has the highest earnings from their books.
SELECT a.author_name, SUM(b.price)
FROM Order_Details od
JOIN Books b ON od.book_id = b.book_id
JOIN Authors a ON b.author_id = a.author_id
GROUP BY a.author_name
ORDER BY SUM(b.price) DESC
LIMIT 1
