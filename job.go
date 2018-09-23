package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// Execute runs query against the database and tracks timing and errors
func Execute(db *sql.DB, index int, query string) ResponseResult {
	var startTime = time.Now()
	_, err := db.Exec(query)
	elapsedMs := int(time.Since(startTime) / time.Millisecond)
	var errorMessage = "None"
	if err != nil {
		errorMessage = err.Error()
		fmt.Println(errorMessage)
	}
	return ResponseResult{
		Index:          index,
		ResponseTimeMs: elapsedMs,
		ErrorMessage:   errorMessage,
	}
}
