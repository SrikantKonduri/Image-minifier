package databases

import (
	"database/sql"
	"fmt"
	"image_server/utils"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

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

	// fmt.Println("User ", os.Getenv("DB_USER"))
	// fmt.Println("Pwd ", dbPWD)

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

func GetProductURLS(db *sql.DB, productId int64) ([]string, error) {
	var emptyArray []string

	query := "SELECT product_images FROM Products WHERE product_id=?"
	row := db.QueryRow(query, productId)

	var imagesStr sql.NullString
	err := row.Scan(&imagesStr)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where no rows were returned
			return emptyArray, nil
		}
		return emptyArray, err
	}

	if imagesStr.Valid {
		urls := strings.Split(imagesStr.String, ",")
		fmt.Println("Result: ", urls)
		urls = urls
		return urls, nil
	}
	return emptyArray, nil
}

func GetCompressedURL(db *sql.DB, productId int64) (string, error) {
	var emptyStr string

	query := "SELECT compressed_product_images FROM Products WHERE product_id=?"
	row := db.QueryRow(query, productId)

	var compressed_img sql.NullString
	err := row.Scan(&compressed_img)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where no rows were returned
			return emptyStr, nil
		}
		return emptyStr, err
	}

	if compressed_img.Valid {
		return compressed_img.String, nil
	}
	return emptyStr, nil
}

func UpdateCompressedURL(db *sql.DB, productId int64, url string) error {
	query := "UPDATE Products SET compressed_product_images = ? WHERE product_id = ?"
	_, err := db.Exec(query, url, productId)
	if err != nil {
		return err
	}
	return nil
}
