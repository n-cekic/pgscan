package main

import (
	"database/sql"
	"encoding/json"
	"errors"
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
	var TestingTable Table
	data, err := json.Marshal(table)
	if err != nil {
		log.Fatal("josn marshalling returned: ", err.Error())
	}
	err = TestingTable.UnmarshalJSON(data)
	if err != nil {
		log.Fatal("josn unmarshalling returned: ", err.Error())
	}
	fmt.Printf("%+v", TestingTable)
	// nicePrintTable(table, columnNames, len(columnNames))
}

type Table []Testing

type Testing struct {
	ID         int8            `json:"id"`
	CreatedAt  time.Time       `json:"created_at"`
	CreatedBy  string          `json:"created_by"`
	Data       json.RawMessage `json:"data"`
	Checked    bool            `json:"checked"`
	ListString []string        `json:"list_string"`
	Col7       []float32       `json:"column_7"`
}

func (t *Testing) UnmarshalJSON(bytes []byte) error {
	// Define a temporary alias to avoid recursion
	type Alias Testing
	aux := &struct {
		CreatedAt string `json:"created_at"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	// Unmarshal into the temporary struct
	if err := json.Unmarshal(bytes, &aux); err != nil {
		return err
	}

	// Parse CreatedAt timestamp
	parsedTime, err := time.Parse(time.RFC3339, aux.CreatedAt)
	if err != nil {
		return errors.New("invalid timestamp format for 'created_at'")
	}
	t.CreatedAt = parsedTime

	return nil
}

func (t *Table) UnmarshalJSON(bytes []byte) error {
	var raw []json.RawMessage
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	var temp Table
	for _, r := range raw {
		var test Testing
		if err := json.Unmarshal(r, &test); err != nil {
			return err
		}
		temp = append(temp, test)
	}

	*t = temp
	return nil
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
