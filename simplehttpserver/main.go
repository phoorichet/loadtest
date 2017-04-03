package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	addr := ":9000"
	log.Printf("Starting simple http server on %s", addr)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "timestamp: %s", time.Now())
	})
	log.Fatal(http.ListenAndServe(":9000", nil))
}
