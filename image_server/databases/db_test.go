package databases_test

import (
	"database/sql"
	"testing"

	"image_server/databases"

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

func TestGetProductURLS(t *testing.T) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "admin",
		Net:    "tcp",
		Addr:   "localhost" + ":" + "3306",
		DBName: "image_minifier",
	}

	dataSourceName := cfg.FormatDSN()

	// Get a database handle
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

	urls, err := databases.GetProductURLS(db, 1)
	urls = urls

	if err != nil {
		t.Errorf("GetProductURLS failed with error: %v", err)
	}
	validity := true
	// fmt.Println(urls[0])
	// fmt.Println(urls[1])
	// fmt.Println(urls[2])
	validity = validity && urls[0] == "https://tinyjpg.com/images/social/website.jpg"
	validity = validity && urls[1] == "https://fdn2.gsmarena.com/vv/pics/vivo/vivo-v20-2.jpg"
	validity = validity && urls[2] == "https://akm-img-a-in.tosshub.com/indiatoday/images/story/202310/motorola-edge-40-neo-170517173-3x4.png?VersionId=cbxSR9.ZchTQh0vHxZJrqJTsp80m9yef"
	if validity == false {
		t.Errorf("Fetched invalid URLS")
	}

}
