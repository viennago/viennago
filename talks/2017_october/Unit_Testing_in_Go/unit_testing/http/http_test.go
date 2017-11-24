// Author(s): Michael Koeppl

package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/calmandniceperson/vienna.go-unit-testing/db"
	"github.com/stretchr/testify/assert"
)

// TestGetSessionCookie makes sure that our GetSessionCookie handler returns
// a cookie with the response.
func TestGetSessionCookie(t *testing.T) {
	testdb := TestLibraryDB{}
	handler := Handler{db: &testdb}

	req := httptest.NewRequest(
		"GET",
		"/session/get",
		nil,
	)
	rec := httptest.NewRecorder()

	handler.GetSessionCookie(rec, req)
	cookieResponse := &http.Response{Header: rec.Header()}

	cookie := cookieResponse.Cookies()[0]
	assert.NotNil(t, cookie)
	assert.Equal(t, "Hello", cookie.Value)
}

// TestLibraryDB is an implementation of the Datastore interface required by
// our Handler structure in the http package. This can be used to replace the
// "real" LibraryDB (db package).
type TestLibraryDB struct{}

// AddBook is a mock function for the TestLibraryDB structure.
func (testdb *TestLibraryDB) AddBook(book *db.Book) error {
	return nil
}

func TestAddBookHandler(t *testing.T) {
	testdb := TestLibraryDB{}
	handler := Handler{&testdb}

	// Get the cookie from the GetSessionCookie handler. This cookie can then
	// be appended to the request to AddBookHandler.
	cookieReq := httptest.NewRequest(
		"GET",
		"/session/get",
		nil,
	)
	cookieRec := httptest.NewRecorder()
	handler.GetSessionCookie(cookieRec, cookieReq)
	cookieResp := &http.Response{Header: cookieRec.Header()}
	cookie := cookieResp.Cookies()[0] // Here is our cookie! It is added further down

	data := db.Book{
		ISBN:   "4213213029801",
		Author: "Stephen King",
		Title:  "IT",
		Pages:  522,
	}

	json, _ := json.Marshal(&data)

	req := httptest.NewRequest(
		"POST",
		"/book/add",
		strings.NewReader(string(json)),
	)
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	handler.AddBookHandler(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
