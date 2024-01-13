package handlers

import (
	"log"
	"net/http"
)

// Hello implements log
type Hello struct {
	l *log.Logger
}

// NewHello allocates and returns a new Hello Handler.
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	// Printing to the console
	h.l.Println("Hello from server handler")
	// Writing data to the response
	_, err := rw.Write([]byte("Hello to the server\n"))
	if err != nil {
		h.l.Println("Error during hello response write", err)
	}
}
