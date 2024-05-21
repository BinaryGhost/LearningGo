package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

// COOL Websites:
//		> https://pkg.go.dev/github.com/ncruces/go-sqlite3/driver#SQLite.OpenConnector
//		> https://pkg.go.dev/database/sql

var (
	db *sql.DB // It is null, but dont know how it could be fixed
)

func init() {
	if err := Connection(db); err != nil {
		log.Fatal(err)
	}
	if err := TableCreate(db); err != nil {
		log.Fatal(err)
	}
}

func Connection(db *sql.DB) error {
	_, err := sql.Open("sqlite3", "./helpers/local.db")
	if err != nil {
		fmt.Println("Connection to DB:")
		log.Fatal(err)
	}
	// Be sure to close everything in main.go
	return nil
}

// Again for testing purposes -> After programm exists delete the previous DB for
// keeping things clean
func TableCreate(db *sql.DB) error {
	/* if db == nil {
		fmt.Println("Not properly initialized")
		os.Exit(1)
	} */

	_, err := db.Exec(`
		DROP TABLE IF EXISTS urlMap;
		CREATE TABLE urlMap (
			id         INTEGER PRIMARY KEY AUTOINCREMENT,
			shortURL   VARCHAR(128) NOT NULL,
			orgnURL    VARCHAR(255) NOT NULL,
		);
	`)

	if err != nil {
		fmt.Println("TableCreate():")
		log.Fatal(err)
	}
	return nil
}

func CreateData(db *sql.DB, shortURL, originURL string) error {
	/* if db == nil {
		fmt.Println("Not properly initialized")
		os.Exit(1)
	} */

	// If the the original URL appears more than once, this will check it
	var count int

	err := db.QueryRow("SELECT COUNT(orgnURL) FROM urlMap WHERE orgnURL = ? ", originURL).Scan(&count)
	if err != nil {
		fmt.Println("CreateData():")
		log.Fatal(err)
	}

	if count > 0 {
		fmt.Println("CreateData(): This URL is already in your Database-")
		os.Exit(1)
	}

	_, err = db.Exec("INSERT INTO urlMap (shortURL, orgnURL) VALUES (?,?);", shortURL, originURL)
	if err != nil {
		fmt.Println("CreateData():")
		log.Fatal(err)
	}
	return nil
}

func ReadShortURL(db *sql.DB, shortURL string) *sql.Row {
	/* 	if db == nil {
		fmt.Println("Not properly initialized")
		os.Exit(1)
	} */

	var orgnURL string

	row := db.QueryRow("SELECT orgnURL FROM urlMap WHERE shortURL = ?", shortURL)

	if err := row.Scan(&orgnURL); err == sql.ErrNoRows {
		fmt.Println("ReadShortURL: DB is empty and cannot be read")
		os.Exit(1)
	} else if err != nil {
		fmt.Println("ReadShortURL():")
		log.Fatal(err)
	}
	return row
}

func DeleteRow(db *sql.DB, originURL string) error {
	/* if db == nil {
		fmt.Println("Not properly initialized")
		os.Exit(1)
	} */

	res, err := db.Exec("DELETE FROM urlMap WHERE orgnURL = ?", originURL)
	if err != nil {
		fmt.Println("DeleteRow():")
		log.Fatal(err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("DeleteRow(): Looking if table was empty or not")
		log.Fatal(err)
	}
	if rowsAffected == 0 {
		fmt.Println("DeleteRow(): The DB is empty or the URL does not exist in it")
		os.Exit(1)
	}
	return nil
}

func ShowAll(db *sql.DB) *sql.Row {
	/* if db == nil {
		fmt.Println("Not properly initialized")
		os.Exit(1)
	}
	*/

	var orgnURL string

	all := db.QueryRow("SELECT * FROM urlMap")

	if err := all.Scan(&orgnURL); err == sql.ErrNoRows {
		fmt.Println("ReadShortURL: DB is empty and cannot be read")
		os.Exit(1)
	} else if err != nil {
		fmt.Println("ReadShortURL():")
		log.Fatal(err)
	}

	return all
}
