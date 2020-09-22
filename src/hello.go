package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("starting Web Server") // (1)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) { // (2)
		fmt.Fprintln(w, "Hello World!!")
	})

	err := http.ListenAndServe(":8080", nil) // (3)
	if err != nil {                          // (4)
		log.Fatalln(err)
	}
}
