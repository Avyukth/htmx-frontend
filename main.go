package main

import (
	"fmt"
	"time"

	"log"
	"net/http"
	"text/template"
)


type Film struct {
	Title string
	Director string
}


func main() {
	fmt.Println("Hello, World!")
	h1:= func(w http.ResponseWriter, _ *http.Request) {
		
		tmpl:= template.Must(template.ParseFiles("index.html"))
		Films:=map[string][]Film{
		"Films":{
			{Title: "The Shawshank Redemption",
            Director: "Subhrajit"},
			{Title: "The Godfather",
			Director: "Subhrajit"},
            {Title: "The Godfather: Part II",
			Director: "Subhrajit"},
            {Title: "The Dark Knight",
			Director: "Nolan"},
            
        },


	}
	
		tmpl.Execute(w,Films)
		
	}

	// handler function #2 - returns the template block with the newly added film, as an HTMX response
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		// tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	// define handlers
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8090", nil))

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
