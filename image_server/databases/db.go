package databases

import (
	"database/sql"
	"fmt"
	"image_server/utils"
	"os"

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
