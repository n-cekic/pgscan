package pgscan

import (
	"database/sql"
)

type AAAA struct {
	DB *sql.DB
}

// Scan is a function that accepts table name and it's columns as arguments
// and returns the result of query.
//
// data is scanned from the DB into a list of interfaces
func (db *AAAA) Scan(tableName string, columnNames []string) ([]map[string]interface{}, error) {
	noOfColumns := len(columnNames)
	table := make([]map[string]interface{}, 0)

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
		columns := make([]interface{}, len(columnNames))
		scanArgs := make([]interface{}, len(columnNames))
		for i := range columns {
			scanArgs[i] = &columns[i]
		}

		if err := rows.Scan(scanArgs...); err != nil {
			return nil, err
		}

		colTypes, _ := rows.ColumnTypes()

		// Create a map for the row
		row := make(map[string]interface{})
		for i, colName := range columnNames {
			row[colName] = columns[i]
		}

		for i, ct := range colTypes {
			tp := ct.ScanType().
			row[ct.Name()] = row[ct.Name()].(tp.)
		}

		table = append(table, row)
	}
	return table, nil
}
