package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":http", http.FileServer(http.Dir("/public")))
}
