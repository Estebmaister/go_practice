package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Data implements log
type Data struct {
	l *log.Logger
}

// NewData allocates and returns a new Data Handler.
func NewData(l *log.Logger) *Data {
	return &Data{l}
}

// ServeHTTP implements the go http.Handler interface
// which dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (d *Data) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Read from the response body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		d.l.Println("Error reading body", err)

		http.Error(rw, "Oops, something goes wrong when reading the request body", http.StatusBadRequest)
		return
	}
	// Printing to the console
	d.l.Printf("Data string from handler: %s\n", data)
	// Writing data to the response
	fmt.Fprintf(rw, "Data '%s'\n", data)
}
