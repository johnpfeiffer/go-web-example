package main

import (
	"database/sql"
	"time"
)

// Note contains text
type Note struct {
	// TODO: use UUID4 instead of integers for ids
	ID      int       `json:"id"`
	Text    string    `json:"note"`
	Created time.Time `json:"created"`
}

// getNotes queries the database and returns all Notes
func getNotes(db *sql.DB) ([]Note, error) {
	var notes []Note
	rows, err := db.Query("SELECT id,note,created FROM notes")
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

// createNote saves a note to the database
func createNote(db *sql.DB, n Note) (int, error) {
	var id int
	// use the postgres extension of RETURNING
	err := db.QueryRow(`INSERT INTO notes (note) VALUES($1) RETURNING id`,
		n.Text).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}
