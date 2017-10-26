// Author(s): Michael Koeppl

package http

import (
	"encoding/json"
	"net/http"

	"github.com/calmandniceperson/vienna.go-unit-testing/db"
)

type Handler struct {
	// IMPORTANT PART HERE!
	// The fact that this is of type Datastore is what allows us to replace
	// this field with our test db.
	db db.Datastore
}

// GetSessionCookie is a http handler that does nothing else than adding a
// cookie named "librarytoken" to the response so our test user can use it.
func (h *Handler) GetSessionCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{Name: "librarytoken", Value: "Hello"}
	http.SetCookie(w, &c)
}

// AddBookHandler "adds" a new book to our "database". It only does so if there
// is a "librarytoken" cookie with a valid value included in the request.
// This is for the test case to show how to test a route with cookies.
func (h *Handler) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("librarytoken")
	if err != nil || cookie == nil || cookie.Value == "" {
		http.Error(w, "No valid cookie provided", http.StatusForbidden)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var data db.Book
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	if data.ISBN == "" || data.Author == "" || data.Title == "" ||
		data.Pages <= 0 {
		http.Error(w, "Missing or invalid fields", http.StatusBadRequest)
		return
	}

	if err := h.db.AddBook(&data); err != nil {
		http.Error(w, "Error adding book to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
