# Image-minifier

### API
The Product server API receives the following data in json format for adding product

- user_id
- product_name
- product_description (text)
- product_images (Array of image URLs)
- product_price (Number)

### Database Schema

#### Users Table

| Column     | Type          |
|------------|---------------|
| id         | Integer       |
| name       | Varchar(255)  |
| mobile     | Varchar(20)  |
| latitude   | Varchar(20)  |
| longitude  | Varchar(20)  |
| created_at | Datetime      |
| updated_at | Datetime      |

#### Products Table

| Column                     | Type          |
|----------------------------|---------------|
| product_id                 | Integer       |
| product_name               | Varchar(255)  |
| product_description        | text          |
| product_images             | text          |
| product_price              | decimal(10,2) |
| compressed_product_images  | text          |
| created_at                 | Datetime      |
| updated_at                 | Datetime      |


---
### Setup

- Make sure you have rabbitMQ, MySQL server, Go installed in your system.
- Set environment variables in both image_server and product_server directories accordingly in .env file.
- Navigate to image_server and product_server directories and run `go mod tidy` to make sure all dependencies are installed in the system.
- Run commands mentioned in `init.sql`

---

### Instructions to run the project

1. To start product server, open terminal and go to `product_server` directory and run `go run main.go`
2. To start image server, open terminal and go to `image_server` directory and run `go run main.go`

---

### Testing

#### Unit Testing

1. Open terminal and navigate to `product_server` , run `go test ./..`
2. Open terminal and navigate to `image_server` , run `go test ./..`

#### Integration Testing

1. Open terminal and navigate to `tests`, run `go test integration_test.go`

---