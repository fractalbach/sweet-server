package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fractalbach/sweet-server/backend/storage"

	_ "github.com/lib/pq"
)

const version = "0.1"

var store storage.Storage

func main() {
	log.Println("Starting Sweet Server, version ", version)
	store.Init()
	// r := mux.NewRouter()
	// r.HandleFunc("/", myHandler)
	http.HandleFunc("/", myHandler)
	log.Print("Listening 8000")
	// log.Fatal(http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r)))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func myHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "You are at url: %s", r.URL.Path)

	// rows, err := db.Query("SELECT title FROM blog")
	// if err != nil {
	// 	w.WriteHeader(500)
	// 	return
	// }
	// var titles []string
	// for rows.Next() {
	// 	var title string
	// 	err = rows.Scan(&title)
	// 	titles = append(titles, title)
	// }
	// json.NewEncoder(w).Encode(titles)
}
