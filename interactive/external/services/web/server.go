package main

import (
	"fmt"
	"io"
	"net/http"
)

// getDescription grabs a description from the describer service that's running
// on the Docker Desktop host system.
func getDescription() (string, error) {
	// Perform an HTTP request (and defer closure of the body stream.)
	response, err := http.Get("http://host.docker.internal:8081/")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Extract the description bytes.
	description, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// Success.
	return string(description), nil
}

func main() {
	// Create an HTTP endpoint that returns a random message.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		description, err := getDescription()
		if err != nil {
			http.Error(w, "unable to grab description", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf(
			"<!doctype html><html><body><h1>You are %s</h1></body></html>",
			description,
		)))
	})

	// Listen and serve requests.
	http.ListenAndServe(":http", nil)
}
