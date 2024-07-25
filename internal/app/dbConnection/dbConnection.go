package dbConnection

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "shortlink"
)

var db *sql.DB

func GetConnection() {
	db = SetupDB()
	fmt.Println("Successfully connected!")
}

func CloseConnection() {
	db.Close()
	fmt.Println("Connection is closed!")
}

func SetupDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	sqldb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = sqldb.Ping()
	if err != nil {
		panic(err)
	}
	return sqldb
}

func AddUrl(originalUrl, shortUrl string) error {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM urls WHERE short_url=$1 AND original_url=$2)`
	err := db.QueryRow(query, shortUrl, originalUrl).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		query = `INSERT INTO urls (short_url, original_url) VALUES ($1, $2)`
		_, err := db.Exec(query, shortUrl, originalUrl)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetUrl(c echo.Context, shortID string, originalURL *string) error {
	query := "SELECT original_url FROM urls WHERE short_url = $1"
	err := db.QueryRow(query, shortID).Scan(&originalURL)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return nil
}

func CheckURLExists(originalURL string) (error, bool) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM urls WHERE original_url=$1)`
	err := db.QueryRow(query, originalURL).Scan(&exists)

	if err != nil {
		return err, exists
	}

	fmt.Println(exists)

	return nil, exists
}

func GetShortURL(originalURL string) (error, string) {
	var shortURL string

	query := "SELECT short_url FROM urls WHERE original_url = $1"
	err := db.QueryRow(query, originalURL).Scan(&shortURL)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("error: URL not found!")
			return err, ""
		}

		return err, ""
	}

	return nil, shortURL
}
