// It’s always a best practice to serve static resources,
// such as .js, .css, and images from the filesystem, or any content delivery network
// (CDN), such as Akamai or Amazon CloudFront, rather than serving it from the web
// Server to avoid overloading the server.

// In this recipe, we are going to serve static files over HTTP using Gorilla Mux
//The static files are contained in static/css/main.css


package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type Person struct {
	Name string
	Id   string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{Id: "21", Name: "Foo"}
	parsedTemplate, _ := template.ParseFiles("templates/first-template.html")
	err := parsedTemplate.Execute(w, person)
	if err != nil {
		log.Printf("Error occurred while executing the templateor writing its output : ", err)
		return
	}
}
func main() {
    // instantiate the gorilla/mux router by calling the
    // NewRouter() handler of the mux router
    router := mux.NewRouter()
	router.HandleFunc("/", renderTemplate).Methods("GET")

    // This returns a handler
    // that serves HTTP requests by removing /static from the request URL's path and
    // invoking the file server. StripPrefix handles a request for a path that doesn't begin
    // with a prefix by replying with an HTTP 404
    router.PathPrefix("/").Handler(http.StripPrefix("/static",http.FileServer(http.Dir("static/"))))
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
