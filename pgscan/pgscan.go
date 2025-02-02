package pgscan

import (
	"database/sql"
	"fmt"
	"log"
)

type AAAA struct {
	DB *sql.DB
}

// Scan is a function that accepts table name and it's columns as arguments
// and returns the result of query.
//
// data is scanned from the DB into a list of interfaces
func (db *AAAA) Scan(tableName string, columnNames []string) [][]interface{} {
	noOfColumns := len(columnNames)
	table := make([][]interface{}, noOfColumns)

	query := "SELECT "
	// concat the columns into a query string
	for i, column := range columnNames {
		query += `"` + column + `"`
		// add a comma if it's not the last column
		if i < noOfColumns-1 {
			query += ", "
		}
	}

	query += " FROM " + tableName

	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data queried successfully!\n")
	defer rows.Close()

	for rows.Next() {
		tmpRow := make([]interface{}, noOfColumns)
		for i := range tmpRow {
			tmpRow[i] = new(interface{})
		}
		if err := rows.Scan(tmpRow...); err != nil {
			log.Print("failed scanning row")
		}
		// rows.ColumnTypes()
		table = append(table, tmpRow)
	}
	return table
	// nicePrintTable(table, columnNames, noOfColumns)
	// for _, colName := range columnNames {
	// 	fmt.Printf("%s\t", colName)
	// }
	// fmt.Println()
	// for _, row := range table {
	// 	for _, colVal := range row {
	// 		valPtr := colVal.(*interface{})
	// 		// val := (*valPtr).(interface{})
	// 		switch v := (*valPtr).(type) {
	// 		case string:
	// 			fmt.Printf("%s\t", v)
	// 		case int64, int32, int, float32, float64:
	// 			fmt.Printf("%d\t", v)
	// 		default:
	// 			fmt.Printf("unknown type\t")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()
}
