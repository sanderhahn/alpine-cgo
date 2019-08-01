package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type foo = struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func buildQueryHandler(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`select id, name from foo`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		w.Header().Set("content-type", "application/json")
		e := json.NewEncoder(w)
		e.SetEscapeHTML(false)

		foo := foo{}
		first := true
		w.Write([]byte("["))
		for rows.Next() {
			if !first {
				w.Write([]byte(","))
			}
			err = rows.Scan(&foo.ID, &foo.Name)
			if err != nil {
				log.Fatal(err)
			}
			e.Encode(foo)
			first = false
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte("]"))
	})
}

func main() {
	database, ok := os.LookupEnv("DATABASE")
	if !ok {
		log.Fatal("DATABASE environment variable is required")
	}

	os.Remove(database)
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	drop table if exists foo;
	create table foo (id integer not null primary key, name text);
	insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz');
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	log.Print("Running at http://localhost:" + port)

	http.HandleFunc("/", buildQueryHandler(db))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
