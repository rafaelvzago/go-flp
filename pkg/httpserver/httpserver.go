package httpserver

import (
	"fmt"
	"log"
	"net/http"
)

func Server() {
	dataPath := "/app/data" // Updated path
	fs := http.StripPrefix("/claim/", http.FileServer(http.Dir(dataPath)))
	http.Handle("/claim/", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the server!")
	})

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
