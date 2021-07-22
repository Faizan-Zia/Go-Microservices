package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Faizan-Zia/microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product", log.LstdFlags)
	hello := handlers.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/", hello)
	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	server.ListenAndServe()

}
