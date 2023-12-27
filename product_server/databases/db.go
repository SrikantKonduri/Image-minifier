package databases

import (
	"database/sql"
	"fmt"
	"os"
	"product_server/utils"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Product struct {
	Product_id                int32
	Product_name              string
	Product_description       string
	Product_price             float32
	Product_images            string
	Compressed_product_images string
	Created_at                string
	Updated_at                string
}
type User struct {
	User_name      string
	User_mobile    string
	User_latitude  string
	User_longitude string
	Created_at     string
	Updated_at     string
}

func ConnectDB() *sql.DB {
	var db *sql.DB
	er := godotenv.Load()
	if er != nil {
		utils.FailOnError(er, "Failed to read from .env")
		return nil
	}

	dbUser := os.Getenv("DB_USER")
	dbPWD := os.Getenv("DB_PWD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	fmt.Println("User ", dbUser)
	fmt.Println("Pwd ", dbPWD)

	cfg := mysql.Config{
		User:   dbUser,
		Passwd: dbPWD,
		Net:    "tcp",
		Addr:   dbHost + ":" + dbPort,
		DBName: dbName,
	}

	dataSourceName := cfg.FormatDSN()

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		utils.FailOnError(err, "Error opening connection")
		return nil
	}

	pingErr := db.Ping()
	if pingErr != nil {
		utils.FailOnError(err, "Error opening pinging")
		return nil
	}
	fmt.Println("[+] Connected to DB")
	return db
}

func AddProduct(db *sql.DB, p Product) (int64, error) {
	query := "INSERT INTO Products (product_name,product_description,product_price,product_images,created_at,updated_at) VALUES (?, ?, ?, ?, ?,?)"
	result, err := db.Exec(query, p.Product_name, p.Product_description, p.Product_price, p.Product_images, p.Created_at, p.Updated_at)
	if err != nil {
		// utils.FailOnError(err, "Error in executing SQL query")
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		// utils.FailOnError(err, "Error in getting ID")
		return -1, err
	}
	return id, nil
}

func AddUser(db *sql.DB, u User) (int64, error) {
	query := "INSERT INTO Users (name,mobile,latitude,longitude,created_at,updated_at) VALUES (?, ?, ?, ?, ?,?)"
	result, err := db.Exec(query, u.User_name, u.User_mobile, u.User_latitude, u.User_longitude, u.Created_at, u.Updated_at)
	if err != nil {
		// utils.FailOnError(err, "Error in executing SQL query")
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		// utils.FailOnError(err, "Error in getting ID")
		return -1, err
	}
	return id, nil
}
