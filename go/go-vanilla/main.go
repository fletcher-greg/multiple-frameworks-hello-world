package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("Error %s", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!  This is Go!")
}
