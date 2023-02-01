package handlers

import (
	
	"log"
	"net/http"
)

type DocumentsHandler struct {
	l *log.Logger
}

//Define handlers
var GetDocHandler=  http.HandlerFunc( func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Getting doc!"))
})

var CreateDocHandler= http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Creating doc!"))
})