package main

import (
	_ "embed"
	"net/http"
	"sync/atomic"
	"text/template"
)

//go:embed template.html
var templateHTML []byte

// templateContext is the type used to pass template information.
type templateContext struct {
	// Count is the request number.
	Count uint64
	// Address is the client's address.
	Address string
}

func main() {
	// Parse the HTML template.
	page, err := template.New("page").Parse(string(templateHTML))
	if err != nil {
		panic("unable to parse page template")
	}

	// Track request counts.
	var count uint64

	// Serve requests.
	http.ListenAndServe(":http", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		page.Execute(w, &templateContext{
			Count:   atomic.AddUint64(&count, 1),
			Address: r.RemoteAddr,
		})
	}))
}
