package handlers


import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	
)

//Creating a simple HTTP server
const
(
CONN_HOST = "localhost"
CONN_PORT = "8080"
)


//Welcome() - takes ResponseWriter and Request as an input and writes
//Welcome to docup! on an HTTP response stream.
func welcome(w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w, "Welcome to docup!")
}

//HTTP request routing


 func main(){
  router:=mux.NewRouter()

   
   router.Handle("/getdocs",GetDocHandler).Methods("GET")
   router.Handle("/createdoc",CreateDocHandler).Methods("POST")
  
   err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
   if err != nil{
   log.Fatal("error starting http server : ", err)
   return
   }
  
 }