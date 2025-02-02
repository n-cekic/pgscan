package main

import (
	"database/sql"
	"fmt"
	"log"
	"pgscan/pgscan"

	_ "github.com/lib/pq"
	tbl "github.com/rodaine/table"
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

	database := pgscan.AAAA{DB: db}
	columnNames := []string{"name", "population"}
	table := database.Scan("city", columnNames)
	nicePrintTable(table, columnNames, len(columnNames))
}

func nicePrintTable(table [][]interface{}, columnNames []string, noOfColumns int) {
	interfaceSlice := make([]interface{}, noOfColumns)
	for i, v := range columnNames {
		interfaceSlice[i] = v
	}
	formatedTable := tbl.New(interfaceSlice...)

	for _, row := range table {
		tmpRow := make([]interface{}, 0, 2)
		for _, colVal := range row {
			valPtr := colVal.(*interface{})
			tmpRow = append(tmpRow, *valPtr)
		}
		formatedTable.AddRow(tmpRow...)

	}

	formatedTable.Print()
}
