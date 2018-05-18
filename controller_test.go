package main

// integration tests that require a database

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetIndex(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	IndexHandler(w, req)
	response := w.Result()
	assertResponseCode(t, http.StatusOK, response.StatusCode)
	body, _ := ioutil.ReadAll(response.Body)
	assertResponseBody(t, defaultIndexResponse, string(body))
}

func TestIntegrationGetIndex(t *testing.T) {
	if !isIntegrationTest() {
		t.Skip()
	}
	// GIVEN
	cleanStart(t)

	// WHEN
	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(t, req)

	// THEN
	assertResponseCode(t, http.StatusOK, response.Code)
	body, _ := ioutil.ReadAll(response.Body)
	assertResponseBody(t, defaultIndexResponse, string(body))

	// CLEANUP
	err := resetTable(t, testDB, NoteTable, NoteTableSequence)
	exitIfError(err)
}

func TestIntegrationPostNote(t *testing.T) {
	if !isIntegrationTest() {
		t.Skip()
	}
	// GIVEN
	cleanStart(t)

	testNoteString := "this is a test note"
	testNoteJSON := `{"note": "` + testNoteString + `"}`
	payload := []byte(testNoteJSON)
	req, _ := http.NewRequest("POST", "/note", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// WHEN
	response := executeRequest(t, req)

	// THEN
	assertResponseCode(t, http.StatusOK, response.Code)

	// WHEN
	reqGet, _ := http.NewRequest("GET", "/note", nil)
	responseGet := executeRequest(t, reqGet)

	// THEN
	assertResponseCode(t, http.StatusOK, response.Code)
	assertResponseCode(t, http.StatusOK, responseGet.Code)
	body, _ := ioutil.ReadAll(responseGet.Body)

	now := time.Now().UTC()
	// If the tests runs exactly at the boundary of an hour this might fail
	timestamp := fmt.Sprintf("%d-%02d-%02dT%02d:", now.Year(), now.Month(), now.Day(), now.Hour())
	expectedJSON := `[
  {
    "id": 1,
    "note": "` + testNoteString + `",
    "created": "` + timestamp
	assertResponseBodyStartsWith(t, expectedJSON, string(body))

	// CLEANUP
	err := resetTable(t, testDB, NoteTable, NoteTableSequence)
	exitIfError(err)
}

func executeRequest(t *testing.T, req *http.Request) *httptest.ResponseRecorder {
	t.Helper()
	w := httptest.NewRecorder()
	// TODO: is there some way to avoid this global test variable
	r := router(testDB)
	r.ServeHTTP(w, req)
	return w
}

func assertResponseCode(t *testing.T, expected, result int) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected response code: %d ,Received: %d\n", expected, result)
	}
}
func assertResponseBodyStartsWith(t *testing.T, expected, result string) {
	t.Helper()
	if strings.HasPrefix(result, expected) {
		return
	}
	t.Errorf("Expected response body starts with: \n%s\nInstead received: \n%s\n", expected, result)
}

func assertResponseBody(t *testing.T, expected, result string) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected response body: \n%s\nInstead received: \n%s\n", expected, result)
	}
}
