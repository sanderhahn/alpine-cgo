package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	log.Print("Running at " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
