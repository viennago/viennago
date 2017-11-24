// Author(s): MIchael Koeppl

package db

import "github.com/jmoiron/sqlx"

type Book struct {
	ISBN   string
	Author string
	Title  string
	Pages  int
}

// Datastore is the interface implemented by all our db structs (LibraryDB,
// TestLibraryDB). This interface is the type of Handler's db field.
type Datastore interface {
	AddBook(*Book) error
}

// LibraryDB is our "real" db in this example. It wraps sqlx.DB to look real.
type LibraryDB struct {
	*sqlx.DB
}

func (db *LibraryDB) AddBook(newBook *Book) error {
	// For this example, we don't really access a database
	return nil
}
