package main

import (
	"log"
	"math/rand"
	"net/http"
)

// descriptions are the descriptions that this service can return.
var descriptions = []string{
	"awesome 🤩",
	"goofy 🙃",
	"creative 🤔",
	"an awesome/goofy/creative developer 🐳",
}

func main() {
	http.ListenAndServe(":8081", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request from", r.RemoteAddr)
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(descriptions[rand.Intn(len(descriptions))]))
	}))
}
