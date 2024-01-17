package main

import (
	"cloud/utils"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	if err := utils.LoadEnv(".env"); err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// prod parse all, dev parse on request
	tmpl := template.Must(template.ParseFiles("./templates/views/index.html", "./templates/components/head.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, map[string]interface{}{
			"Name": "Cloud",
		})
		if err != nil {
			return
		}
	})

	fmt.Print("Server running on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Error starting the server")
		return
	}
}
