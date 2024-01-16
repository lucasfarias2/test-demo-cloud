package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World test"))
		if err != nil {
			log.Println("Error writing response")
			return
		}
	})

	log.Println("Server is running")

	err := http.ListenAndServe(":8080", nil)git 
	if err != nil {
		log.Fatalln("Error starting the server")
		return
	}
}
