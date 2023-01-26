//Creating and reading HTML forms

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

//User struct has Username and Password fields
//which correspond to the values the login-form.html expects
type User struct
{
Username string
Firstname string
Password string
}
// Here we define a login handler, which checks if the HTTP request calling the handler
// is a GET request and then parses login-form.html from the templates directory
// and writes it to an HTTP response stream; otherwise, it calls the readForm handler
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method=="GET"{parsedTemplate, _ := template.ParseFiles("templates/login-form.html")
	parsedTemplate.Execute(w, nil)
}else{
user := readForm(r)
fmt.Fprintf(w, "Hello "+user.Username+"!"+ "your first name is "+user.Firstname)
}
	
}
func readForm(r *http.Request) *User{
// 	Here we parse the request body as a form and put the results into
// both r.PostForm and r.Form.
r.ParseForm()
user := new(User) //Here we create a new User struct type.
// Here we decode parsed form data from
// POST body parameters to a user struct.
decoder := schema.NewDecoder()
decodeErr := decoder.Decode(user, r.PostForm)
if decodeErr != nil{
log.Printf("error mapping parsed form data to struct : ",
decodeErr)
}
return user}

func main() {
    router := mux.NewRouter()
    // Here we are registering a login function with the / URL
    //pattern using HandleFunc of the gorilla/mux package
	router.HandleFunc("/login", login)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
