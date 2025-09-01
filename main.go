package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/Hello", helloHandler)
	mux.HandleFunc("/mario", marioHandler)

	log.Println("Starting on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Ini landing page"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello dunia, aku lagi belajar golank"))
}

func marioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario from nitendo"))
}