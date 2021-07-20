package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oh no!", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s", d)
		log.Println("Get Requests")
	})

	http.ListenAndServe(":9090", nil)
}
