package pgscan

import (
	"database/sql"
	"time"
)

type AAAA struct {
	DB *sql.DB
}

// Scan is a function that accepts table name and it's columns as arguments
// and returns the result of query.
//
// data is scanned from the DB into a list of interfaces
func (db *AAAA) Scan(tableName string, columnNames []string) ([][]interface{}, error) {
	noOfColumns := len(columnNames)
	table := make([][]interface{}, 0, noOfColumns)

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
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tmpRow := make([]interface{}, noOfColumns)
		for i := range tmpRow {
			tmpRow[i] = new(interface{})
		}
		if err := rows.Scan(tmpRow...); err != nil {
			return nil, err
		}
		for i := range tmpRow {
			a, _ := tmpRow[i].(*interface{})

			switch x := (*a).(type) {
			case string:
				tmpRow[i] = x
			case int, int8, int16, int32, int64:
				tmpRow[i] = x
			case float32, float64:
				tmpRow[i] = x
			case bool:
				tmpRow[i] = x
			case time.Time:
				tmpRow[i] = x
			case []byte:
				tmpRow[i] = x
			}
		}
		table = append(table, tmpRow)
	}
	return table, nil
}
