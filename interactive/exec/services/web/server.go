package main

import (
	"net/http"
)

func main() {
	// Define a simple handler.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Hello, world!"))
	})

	// Listen and serve requests.
	http.ListenAndServe(":http", nil)
}
