package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Faizan-Zia/microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product", log.LstdFlags)
	product := handlers.NewProducts(l)
	sm := http.NewServeMux()
	sm.Handle("/", product)
	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	l.Println("Listening on port 9090")
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println("Gracefully shutting down the server", sig)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
