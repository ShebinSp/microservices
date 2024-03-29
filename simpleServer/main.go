package main

import (
	"context"
	"log"
	"microServices/simpleServer/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
) 

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	// create handlers
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)

	// create a new serve mux and registe the handlers
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/exit", gb)

	// create a new server
	s := &http.Server{
		Addr:         ":8090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// start the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Received terminate, graceful Shutdown",sig)
	
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	s.Shutdown(ctx)
	defer cancel()
}
