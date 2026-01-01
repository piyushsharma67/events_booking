package database

import (
	"database/sql"
	_ "embed"
)

//go:embed schema.sql
var schemaSQL string

func initSchema(db *sql.DB) error {
	if schemaSQL == "" {
        panic("schemaSQL is empty — embed not working")
    }
    _, err := db.Exec(schemaSQL)
    return err
}
