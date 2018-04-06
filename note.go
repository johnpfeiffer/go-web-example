package main

import (
	"database/sql"
	"time"
)

// Note contains text
type Note struct {
	ID      int       `json:"id"`
	Text    string    `json:"note"`
	Created time.Time `json:"created"`
}

// getNotes queries the database
func getNotes(db *sql.DB) ([]Note, error) {
	var notes []Note
	rows, err := db.Query("SELECT id,note,created FROM Notes")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		var text string
		var created time.Time
		err = rows.Scan(&id, &text, &created)
		if err != nil {
			return nil, err
		}
		notes = append(notes, Note{ID: id, Text: text, Created: created})
	}
	return notes, nil
}
