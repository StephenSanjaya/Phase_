CREATE TABLE users (
  user_id SERIAL PRIMARY KEY, 
  nama VARCHAR (50) NOT NULL, 
  status VARCHAR (50) NOT NULL, 
  deposit_amount INTEGER
);