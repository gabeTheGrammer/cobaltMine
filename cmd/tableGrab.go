package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type TableGet struct {
	DB *sql.DB
}

// Structure of request for a table sent by JS
type tableRequest struct {
	Table string `json:"table"`
}

func (app *application) tableHelper(w http.ResponseWriter, r *http.Request) {
	var req tableRequest

	// Gets request from JS and decodes it into req and checks if the request is valid
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		app.errLog.Fatal(err)
		return
	}

	// Simple check if request is empty
	if req.Table == "" {
		http.Error(w, "Table name is required", http.StatusBadRequest)
	}

	// Send sql request to DB
	query := fmt.Sprintf("SELECT * FROM `%s`", req.Table)
	rows, err := app.DB.DB.Query(query)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}

	defer rows.Close()

	// Columns and rows for retrived data
	columns, _ := rows.Columns()
	values := make([]any, len(columns))
	valuePtr := make([]any, len(columns))

	allrows := [][]string{}

	for rows.Next() {
		// Create the pointers to the address of the values
		for i := range columns {
			valuePtr[i] = &values[i]
		}

		// Retrieve the row data
		rows.Scan(valuePtr...)

		// Clean up data
		rowData := make([]string, len(columns))
		for i, val := range values {
			if b, ok := val.([]byte); ok {
				rowData[i] = string(b)
			} else if val != nil {
				rowData[i] = fmt.Sprint(val)
			}
		}

		// Add rows to final data
		allrows = append(allrows, rowData)
	}

	// Encodes data to json
	json.NewEncoder(w).Encode(map[string]any{
		"columns": columns,
		"rows":    allrows,
	})
}
