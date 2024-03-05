-- Challenge 1: Product Count by Price Category
SELECT * ,
CASE
	WHEN Price < 300 THEN "LOW PRICE"
    WHEN Price >= 300 AND Price <= 600 THEN "MEDIUM PRICE"
    WHEN Price > 600 THEN "HIGH PRICE"
END AS "Category Price"
FROM products;

-- Challenge 2: Top N Products by Sales
SELECT p.ProductName, ps.QuantitySold
FROM productsales ps
JOIN products p ON ps.ProductID = p.ProductID
ORDER BY ps.QuantitySold DESC
LIMIT 3;

-- Challenge 3: Product Sales Growth
SELECT ps.ProductID, p.ProductName, SUM(CASE WHEN MONTH(ps.SaleDate) = 8 THEN ps.QuantitySold ELSE 0 END) AS CurrentSold, SUM(CASE WHEN MONTH(ps.SaleDate) = 7 THEN ps.QuantitySold ELSE 0 END) AS PreviousSold
FROM productsales ps
JOIN products p ON ps.ProductID = p.ProductID
GROUP BY ps.ProductID

-- Challenge 4: Average Price Difference
SELECT p1.ProductID, p1.ProductName, p1.Price, (p1.Price - p2.Price) AS Average
FROM products p1
INNER JOIN Products p2
WHERE p1.ProductID - p2.ProductID = 1
ORDER BY p1.ProductID

-- Challenge 5: Products without Price
SELECT *
FROM products
WHERE price IS NULL