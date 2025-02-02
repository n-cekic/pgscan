package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"pgscan/pgscan"
	"time"

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

	// columnNames := []string{"name", "population"}
	// table, err := database.Scan("city", columnNames)
	columnNames := []string{"id", "created_at", "created_by", "data", "checked", "list_string", "column_7"}
	table, err := database.Scan("testing", columnNames)
	if err != nil {
		log.Fatal("database scan returned: ", err.Error())
	}
	var TestingTable []Testing
	data, _ := json.Marshal(table)
	json.Unmarshal(data, &TestingTable)
	fmt.Printf("%+v", TestingTable)

	// jsonTable, err := json.Marshal(table)
	// if err != nil {
	// 	log.Fatal("josn marshalling returned: ", err.Error())
	// }
	// err = json.Unmarshal(jsonTable, &TestingTable)
	// if err != nil {
	// 	log.Fatal("josn unmarshalling returned: ", err.Error())
	// }
	// nicePrintTable(table, columnNames, len(columnNames))
}

type Testing struct {
	ID         int8            `json:"id"`
	CreatedAt  time.Time       `json:"created_at"`
	CreatedBy  string          `json:"created_by"`
	Data       json.RawMessage `json:"data"`
	Checked    bool            `json:"checked"`
	ListString []string        `json:"list_string"`
	Col7       []float32       `json:"column_7"`
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
