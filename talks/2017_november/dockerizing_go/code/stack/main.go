package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func connectDB() {
	// Doesn't work (replace localhost with service name)!
	connStr := "host=localhost user=gopher password=gopherpw dbname=viennago sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s called: %s\n", r.URL.Path, r.Method)
	var text string
	err := db.QueryRow(
		"SELECT txt FROM viennago.Texts",
	).Scan(&text)
	if err != nil {
		fmt.Printf("Failed to read from DB: %s", err.Error())
	}

	fmt.Fprintf(w, "%s", text)
}

func main() {
	connectDB()
	defer db.Close()
	println("Listening on port :8080")
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
