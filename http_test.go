package main

// integration tests that require a database

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIntegrationGetIndex(t *testing.T) {
	if !isIntegrationTest() {
		t.Skip()
	}
	// GIVEN CLEAN STATE
	err := resetTable(t, testDB, NoteTable, NoteTableSequence)
	exitIfError(err)

	preNotes, err := getNotes(testDB)
	exitIfError(err)
	if len(preNotes) != 0 {
		log.Fatal("No notes should exist yet in the test")
	}

	// WHEN

	// payload := []byte(`{"text": "this is some note text"}`)
	// req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(payload))
	// json.Unmarshal(response.Body.Bytes(), &notesObject)
	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(t, req)

	// THEN
	assertResponseCode(t, http.StatusOK, response.Code)

	body, _ := ioutil.ReadAll(response.Body)
	assertResponseBody(t, defaultIndexResponse, string(body))

	// CLEANUP
	err = resetTable(t, testDB, NoteTable, NoteTableSequence)
	exitIfError(err)
}

// TODO: migrate this snippet to unit testing of controller-index
// IndexHandler(w, req)
// https://golang.org/pkg/net/http/httptest/#example_ResponseRecorder

func executeRequest(t *testing.T, req *http.Request) *httptest.ResponseRecorder {
	t.Helper()
	w := httptest.NewRecorder()
	router().ServeHTTP(w, req)
	return w
}

func assertResponseCode(t *testing.T, expected, result int) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected response code: %d ,Received: %d\n", expected, result)
	}
}

func assertResponseBody(t *testing.T, expected, result string) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected response body: %s ,Received: %s\n", expected, result)
	}
}
