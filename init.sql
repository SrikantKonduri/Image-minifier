CREATE DATABASE IF NOT EXISTS image_minifier;

USE image_minifier;

CREATE TABLE IF NOT EXISTS Users (
   id INT PRIMARY KEY AUTO_INCREMENT,
   name VARCHAR(255),
   mobile VARCHAR(20),
   latitude VARCHAR(20),
   longitude VARCHAR(20),
   created_at DATETIME,
   updated_at DATETIME
);


CREATE TABLE IF NOT EXISTS Products (
   product_id INT PRIMARY KEY AUTO_INCREMENT,
   product_name VARCHAR(255),
   product_description TEXT,
   product_images TEXT,
   product_price DECIMAL(10,2),
   compressed_product_images TEXT,
   created_at DATETIME,
   updated_at DATETIME
);
