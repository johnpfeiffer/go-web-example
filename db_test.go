package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestIntegrationNotesCreate(t *testing.T) {
	if !isIntegrationTest() {
		t.Skip()
	}
	// CLEAN STATE
	err := resetTable(t, testDB, NoteTable, NoteTableSequence)
	exitIfError(err)

	preNotes, err := getNotes(testDB)
	exitIfError(err)
	if len(preNotes) != 0 {
		log.Fatal("No notes should exist yet in the test")
	}

	// WHEN
	testText := "a test note"
	testNote := Note{Text: testText}
	noteID, err := createNote(testDB, testNote)
	assertItemID(t, noteID, 1, err, "A single test note should exist with id 1, instead received id:"+strconv.Itoa(noteID))

	testNote.Text = testNote.Text + " second"
	noteID2, err := createNote(testDB, testNote)
	assertItemID(t, noteID2, 2, err, "A second test note should exist with id 2, instead received id:"+strconv.Itoa(noteID2))

	// THEN
	postNotes, err := getNotes(testDB)
	if len(postNotes) != 2 {
		log.Fatal("Two test notes should exist")
	}
	if postNotes[0].Text != testText {
		log.Fatal("A test note text should be", testText, "but instead received", postNotes[0].Text)
	}

	// CLEANUP
	err = resetTable(t, testDB, NoteTable, NoteTableSequence)
	exitIfError(err)
}

// Helper functions

func assertItemID(t *testing.T, id, expected int, err error, errorMessage string) {
	t.Helper()
	exitIfError(err)
	if id != expected {
		log.Fatal(errorMessage)
	}
}

func resetTable(t *testing.T, db *sql.DB, name, sequenceName string) error {
	t.Helper()
	// Truncate is more efficient but does not work if a table contains foreign key(s)
	command := fmt.Sprintf("DELETE FROM %s", name)
	_, err := db.Exec(command)
	if err != nil {
		return err
	}
	if sequenceName != "" {
		command = fmt.Sprintf("ALTER SEQUENCE %s RESTART WITH 1", sequenceName)
		_, err := db.Exec(command)
		if err != nil {
			return err
		}
	}
	return nil
}
