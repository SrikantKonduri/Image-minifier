# Image-minifier

### API
The Product server API receives the following data in json format for adding product

- user_id
- product_name
- product_description (text)
- product_images (Array of image URLs)
- product_price (Number)

---
### Setup

- Make sure you have rabbitMQ, MySQL server, Go installed in your system.
- Set environment variables in both image_server and product_server directories accordingly in .env file.
- Navigate to image_server and product_server directories and run `go mod tidy` to make sure all dependencies are installed in the system.

---

### Instructions to run the project

1. To start product server, open terminal and go to `product_server` directory and run `go run main.go`
2. To start image server, open terminal and go to `image_server` directory and run `go run main.go`

---
