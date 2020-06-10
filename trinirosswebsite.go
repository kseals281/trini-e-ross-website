package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseGlob("*.html"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("handling index")
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func aboutHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("handling about")
	err := tpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		log.Println(err)
	}
}
