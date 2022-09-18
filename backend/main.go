package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fractalbach/sweet-server/backend/storage"
)

const version = "0.01"

var store storage.Storage

func main() {
	log.Println("Starting Sweet Server, version ", version)
	store.Init()
	addExampleUsers()
	http.HandleFunc("/", myHandler)
	log.Print("Listening 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("You are at url: %s\n%s", r.URL.Path, store.PrintWholeTable())
	fmt.Fprintf(w, s)
}

func addExampleUsers() {
	names := []string{"apple", "orange", "awesomer", "raindeer", "wooo"}
	for i := 0; i < 5; i++ {
		user := storage.User{
			Id:   uint(i),
			Name: names[i],
			Data: "",
		}
		store.AddUser(user)
	}
}
