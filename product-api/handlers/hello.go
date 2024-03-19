package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Hello is a simple handler
type Hello struct {
	l *log.Logger
}

// NewHello creates a new hello handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hellow World")
	rw.Write([]byte("Welcome to the World!"))

	// read the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write([]byte("Ooops an error occured"))

		// http package has an error function which will not stop the flow of the code
		http.Error(rw, "Ooops an error occured", http.StatusBadRequest)
		return
	}
	//	log.Printf("Request is %s",d)
	
	// write the response
	fmt.Fprintf(rw, "\nHei %s\n", body)

}
