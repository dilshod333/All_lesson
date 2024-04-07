package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func check(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		switch r.URL.Path {
		case "/":
			fmt.Fprint(w, "Hi main")
		case "/about/":
			temp2 := template.Must(template.ParseFiles("demo.html"))
			temp2.Execute(w, nil)
		case "/dilshod/":
			temp3 := template.Must(template.ParseFiles("demo.html"))
			temp3.Execute(w, nil)
		}
	}
}

func main() {
	fmt.Println("Server running at 8080 Port: ")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("index.html"))
		temp.Execute(w, nil)
		io.WriteString(w, "Main Page\n")
		io.WriteString(w, r.Method)
		fmt.Fprint(w, r.Method)
	}
	http.HandleFunc("/", check)
	http.HandleFunc("/Salom/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
