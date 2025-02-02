package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "world"
	password = "world123"
	dbname   = "world-db"
)

func main() {
	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

	database := AAAA{db: db}
	database.scan("city", []string{"name", "population"})
}

type AAAA struct {
	db *sql.DB
}

// scan is a function that accepts table name and it's columns as arguments
// and returns the result of query.
//
// data is scanned from the DB into a list of interfaces
func (db *AAAA) scan(tableName string, columns []string) {
	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{}) // Allocate memory for each column
	}
	query := "SELECT "
	// concat the columns into a query string
	for i, column := range columns {
		query += `"` + column + `"`
		// add a comma if it's not the last column
		if i < len(columns)-1 {
			query += ", "
		}
	}

	query += " FROM " + tableName

	err := db.db.QueryRow(query).Scan(values...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data scanned successfully!\n")
	for _, c := range columns {
		fmt.Printf("%s\t", c)
	}
	fmt.Println()
	for _, v := range values {
		s, _ := (v).(*interface{})
		s1, _ := (*s).(interface{})
		switch s1.(type) {
		case string:
			s12, _ := (s1).(string)
			fmt.Printf("%s\t", s12)
			continue
		case int32, int64, int, float32, float64:
			s12, _ := (s1).(int64)
			fmt.Printf("%d\t", s12)
			continue
		}
	}
	fmt.Println()
}
