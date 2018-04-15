package main

import (
	"database/sql"
	"fmt"
	"time"
)

// NoteTable is the table where notes are stored in the database
const NoteTable = "notes"

// NoteTableSequence is the sequential id (primary key) for notes in the database
const NoteTableSequence = "notes_id_seq"

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
	query := fmt.Sprintf("SELECT id,note,created FROM %s", NoteTable)
	rows, err := db.Query(query)
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
	command := fmt.Sprintf("INSERT INTO %s (note) VALUES($1) RETURNING id", NoteTable)
	err := db.QueryRow(command, n.Text).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}
