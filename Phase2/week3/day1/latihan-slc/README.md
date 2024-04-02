# FTGO-P2-V1-SLC2
## RULES
1. **Untuk kampus remote**: **WAJIB** melakukan **share screen**(**DESKTOP/ENTIRE SCREEN**) dan **unmute microphone** ketika Live Code
berjalan (tidak melakukan share screen/salah screen atau tidak unmute microphone akan di ingatkan).
2. Kerjakan secara individu. Segala bentuk kecurangan (mencontek ataupun diskusi) akan menyebabkan skor live code ini 0.
3. Clone repo ini kemudian buatlah **branch dengan nama kalian**.
4. Kerjakan pada file Golang (\*.go) yang telah disediakan.
5. Waktu pengerjaan: **90 menit** untuk **2 soal**.
6. **Pada text editor hanya ada file yang terdapat pada repository ini**.
7. Membuka referensi eksternal seperti Google, StackOverflow, dan MDN diperbolehkan.
8. Dilarang membuka repository di organisasi tugas, baik pada organisasi batch sendiri ataupun batch lain, baik branch sendiri maupun branch orang
lain (**setelah melakukan clone, close tab GitHub pada web browser kalian**).
9. Lakukan `git push origin <branch-name>` dan create pull request **hanya jika waktu Live Code telah usai (bukan ketika kalian sudah selesai
mengerjakan)**. Tuliskan nama lengkap kalian saat membuat pull request dan assign buddy.
10. **Penilaian berbasis logika dan hasil akhir**. Pastikan keduanya sudah benar.




## Notes
Live code ini memiliki bobot nilai sebagai berikut:

|Criteria|Meet Expectations|Points|
|---|---|---|---|---|
|Problem Solving - Entity Design|All necessary entities are included |10 pts |
|   |Entity attributes are clearly defined and appropriate|10 pts |
|Relationship Design |All necessary relationships are defined|10 pts|
||Relationship cardinalities are correct and implemented through junction tables as necessary|10 pts|
|Attribute Design |Attributes are properly defined|10 pts|
||Attribute data types and constraints are correctly implemented|10 pts|
|Normalization|The schema is normalized to at least third normal form|10 pts|
||There are no redundant or duplicate attributes or relationships|10 pts|
|Query Support |The desired queries can be supported by the schema design|5 pts|
||Appropriate indexes and constraints are implemented to optimize query performance |5 pts|
|Database Functionality |Stored procedures, triggers or other advanced database features improve functionality and efficiency |5 pts|
||Database operations are handled efficiently |5 pts|



####KETENTUAN
Here are some dos and don'ts to consider when working on this livecode:

Do:

- Identify all the entities involved in the system, including customers, products, orders, and payments.
- Use clear and concise naming conventions for each entity and attribute.
- Clearly define the relationships between entities, such as the one-to-many relationship between customers and orders, and the many-to-many relationship between orders and products.
- Include foreign keys in the appropriate tables to ensure data integrity and enforce referential integrity constraints.
- Use appropriate data types for each attribute, such as using a date data type for order dates and a decimal data type for prices.
- Use indexes where necessary to improve performance for frequently queried data.

Don't:

- Use vague or ambiguous naming conventions that make it difficult to understand the purpose of each entity and attribute.
- Overcomplicate the relationships between entities, such as creating unnecessary junction tables or introducing circular references.
- Store redundant data in multiple tables, which can lead to data inconsistencies and update anomalies.
- Use inappropriate data types for attributes, such as using a string data type for numerical values.
- Neglect to define primary and foreign keys, which can lead to data integrity issues and orphaned records.
- Neglect to consider the performance implications of our schema, such as the need for indexes on frequently queried data.




# SIM LIVECODE 2
## **E_Commerce ERD**

## Restrictions
-




## Objectives
- Mampu memahami konsep database PostgreSQL
- Mampu membuat Database PostgreSQL From scratch
- Mampu membuat Entity Relationship Diagram



## Directions
Our e-commerce platform allows customers to purchase products online. Each customer can create an account with their personal information such as their name, email, phone number, and address. Customers can browse a catalog of available products and add them to their cart. Once they are ready to checkout, they can place an order which will be recorded in our system.


Each order can contain multiple products, each with their own quantity. When an order is placed, we need to record the order information, the customer who placed the order, and the individual products that were ordered. We also need to record the payment information for the order, such as the payment method and confirmation details.


Our system should allow customers to view their past orders and the details of each order, including the products they ordered, the date the order was placed, the payment method used, and the payment confirmation details. We should also be able to view a catalog of available products, as well as the inventory levels and pricing information for each product.

## DB Requirement Tasks

**A. Customers table** :
   Attributes: customer_id (primary key), first_name, last_name, email, phone_number, address.
   
   Explanation: This table stores information about each customer that has created an account on our e-commerce platform. The customer_id serves as the primary key for the table. There is a one-to-many relationship between customers and orders, as each customer can place multiple orders.


**B. Products table** :
   Attributes: product_id (primary key), name, description, price, inventory_level. 
   
   Explanation: This table stores information about each product that is available for purchase on our e-commerce platform. The product_id serves as the primary key for the table. There is a many-to-many relationship between products and orders, as each order can contain multiple products and each product can be ordered in multiple orders.


**C. Orders table** :
   Attributes: order_id (primary key), order_date, customer_id (foreign key), payment_method, payment_confirmation.
   
   Explanation: This table stores information about each order that is placed on our e-commerce platform. The order_id serves as the primary key for the table. There is a one-to-many relationship between customers and orders, as each customer can place multiple orders. There is also a many-to-many relationship between orders and products, as each order can contain multiple products and each product can be ordered in multiple orders. The customer_id attribute serves as a foreign key referencing the customer_id attribute in the customers table.


**D. Order_Items table**:
   Attributes: order_id (foreign key), product_id (foreign key), quantity.
   
   Explanation: This table represents the many-to-many relationship between orders and products. It stores information about the individual products that are included in each order, as well as the quantity of each product. The order_id attribute serves as a foreign key referencing the order_id attribute in the orders table, and the product_id attribute serves as a foreign key referencing the product_id attribute in the products table.


**E. Payments table**:
   Attributes: payment_id (primary key), order_id (foreign key), amount, payment_date.
   
   Explanation: This table stores information about the payment details for each order that is placed on our e-commerce platform. The payment_id serves as the primary key for the table. There is a one-to-one relationship between orders and payments, as each order should have a corresponding payment. The order_id attribute serves as a foreign key referencing the order_id attribute in the orders table.
