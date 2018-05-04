package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func normal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Second)
	fmt.Fprintf(w, "timeout")
}

func main() {
	http.HandleFunc("/normal", normal)
	http.HandleFunc("/timeout", timeout)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
