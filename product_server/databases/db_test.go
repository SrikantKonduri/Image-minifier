package databases

import (
	"database/sql"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
)

func TestConnectDB(t *testing.T) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "admin",
		Net:    "tcp",
		Addr:   "localhost" + ":" + "3306",
		DBName: "image_minifier",
	}

	dataSourceName := cfg.FormatDSN()

	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		t.Errorf("ConnectDB failed to establish a connection")
	}

	pingErr := db.Ping()
	if pingErr != nil {
		t.Errorf("ConnectDB failed to ping the database: %v", pingErr)
	}
	defer db.Close()
}

func TestAddProduct(t *testing.T) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "admin",
		Net:    "tcp",
		Addr:   "localhost" + ":" + "3306",
		DBName: "image_minifier",
	}

	dataSourceName := cfg.FormatDSN()

	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		t.Errorf("ConnectDB failed to establish a connection")
	}

	pingErr := db.Ping()
	if pingErr != nil {
		t.Errorf("ConnectDB failed to ping the database: %v", pingErr)
	}
	defer db.Close()

	// Create a sample product for testing
	product := Product{
		Product_name:        "Test Product",
		Product_description: "Test Description",
		Product_price:       19.99,
		Product_images:      "image1.jpg,image2.jpg",
		Created_at:          time.Now().Format("2006-01-02 15:04:05"),
		Updated_at:          time.Now().Format("2006-01-02 15:04:05"),
	}

	// Call the AddProduct function
	id, err := AddProduct(db, product)

	// Check if there is no error
	if err != nil {
		t.Errorf("AddProduct failed with error: %v", err)
	}

	// Check if the returned ID is greater than 0
	if id <= 0 {
		t.Errorf("AddProduct returned an invalid ID: %d", id)
	}
}

func TestAddUser(t *testing.T) {
	// Create a mock database connection (you may need to adjust this based on your database setup)
	cfg := mysql.Config{
		User:   "root",
		Passwd: "admin",
		Net:    "tcp",
		Addr:   "localhost" + ":" + "3306",
		DBName: "image_minifier",
	}

	dataSourceName := cfg.FormatDSN()

	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		t.Errorf("ConnectDB failed to establish a connection")
	}

	pingErr := db.Ping()
	if pingErr != nil {
		t.Errorf("ConnectDB failed to ping the database: %v", pingErr)
	}
	defer db.Close()

	// Create a sample user for testing
	user := User{
		User_name:      "Test User",
		User_mobile:    "1234567890",
		User_latitude:  "40.7128",
		User_longitude: "-74.0060",
		Created_at:     time.Now().Format("2006-01-02 15:04:05"),
		Updated_at:     time.Now().Format("2006-01-02 15:04:05"),
	}

	// Call the AddUser function
	id, err := AddUser(db, user)

	// Check if there is no error
	if err != nil {
		t.Errorf("AddUser failed with error: %v", err)
	}

	// Check if the returned ID is greater than 0
	if id <= 0 {
		t.Errorf("AddUser returned an invalid ID: %d", id)
	}
}

func TestVerifyUser(t *testing.T) {
	// Create a mock database connection (you may need to adjust this based on your database setup)
	cfg := mysql.Config{
		User:   "root",
		Passwd: "admin",
		Net:    "tcp",
		Addr:   "localhost" + ":" + "3306",
		DBName: "image_minifier",
	}

	dataSourceName := cfg.FormatDSN()

	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		t.Errorf("ConnectDB failed to establish a connection")
	}

	pingErr := db.Ping()
	if pingErr != nil {
		t.Errorf("ConnectDB failed to ping the database: %v", pingErr)
	}
	defer db.Close()

	// Create a sample user for testing
	user := User{
		User_name:      "Test Verify User",
		User_mobile:    "9876543210",
		User_latitude:  "4.7128",
		User_longitude: "-4.0060",
		Created_at:     time.Now().Format("2006-01-02 15:04:05"),
		Updated_at:     time.Now().Format("2006-01-02 15:04:05"),
	}

	// Call the AddUser function
	id, err := AddUser(db, user)

	// Check if there is no error
	if err != nil {
		t.Errorf("AddUser failed with error: %v", err)
	}

	// Check if the returned ID is greater than 0
	if id <= 0 {
		t.Errorf("AddUser returned an invalid ID: %d", id)
	}

	// Call the VerifyUser function with the sample user ID
	result, err := VerifyUser(db, int(id))

	// Check if there is no error
	if err != nil {
		t.Errorf("VerifyUser failed with error: %v", err)
	}

	// Check if the returned result is true (user exists)
	if !result {
		t.Errorf("VerifyUser returned false for an existing user")
	}

	// Call the VerifyUser function with a non-existent user ID
	result, err = VerifyUser(db, int(id+1))

	// Check if there is no error
	if err != nil {
		t.Errorf("VerifyUser failed with error: %v", err)
	}

	// Check if the returned result is false (user does not exist)
	if result {
		t.Errorf("VerifyUser returned true for a non-existent user")
	}
}
