package sqlprepare

import (
	"database/sql"
	"fmt"
)

// ToPrepare holds a query that should be prepared.
// A valid target must be set to store the prepared stmt.
// An optional name is possible for better error messages.
type ToPrepare struct {
	Query  string
	Name   string
	Target **sql.Stmt
}

// Prepare prepares many SQL Statements for later use.
func Prepare(db *sql.DB, toPrepare ...ToPrepare) error {
	for _, t := range toPrepare {
		stmt, err := db.Prepare(t.Query)
		if err != nil {
			if t.Name != "" {
				return fmt.Errorf("Can't prepare stmt %q: %w", t.Name, err)
			}
			return fmt.Errorf("Can't prepare stmt: %w", err)
		}
		*t.Target = stmt
	}
	return nil
}
