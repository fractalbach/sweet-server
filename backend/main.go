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
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Welcome! Here's a list of all users:\n%s", store.PrintWholeTable())
	} else {
		name := r.URL.Path[1:]
		u, ok := store.GetUserDataByName(name)
		if !ok {
			fmt.Fprintf(w, "The user [%s] does not exist!", name)
		} else {
			fmt.Fprintf(w, "You are at url: %s\nHere is their data:\n%v", name, u)
		}
	}
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
