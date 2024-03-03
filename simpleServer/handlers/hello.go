package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hellow World")
	rw.Write([]byte("Welcome to the World!"))
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write([]byte("Ooops an error occured"))

		// http package has an error function which will not stop the flow of the code
		http.Error(rw, "Ooops an error occured", http.StatusBadRequest)
		return
	}
	//	log.Printf("Request is %s",d)
	fmt.Fprintf(rw, "\nHei %s\n", body)

}
