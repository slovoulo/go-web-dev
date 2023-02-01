// Uploading files to a server
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"html/template"
)



const
(
CONN_HOST = "localhost"
CONN_PORT = "8080"
)


//fileHandler() handler, gets the file from the request, reads its
//content, and eventually writes it onto a file on a server
func fileHandler(w http.ResponseWriter, r *http.Request){
    //Call the FormFile handler on the
   // HTTP request to get the file for the provided form key
file, header, err := r.FormFile("file")

    //Check whether there is any problem while getting the file from the request.
if err != nil{
log.Printf("error getting a file for the provided form key : ",
err)
return
}
//defer statement closes the file once we return from the function
defer file.Close()

// Here we are creating a file named
// uploadedFile inside a /tmp directory(server) with mode 666, which means the client can read
// and write but cannot execute the file
out, pathError := os.Create("/tmp/uploadedFile")

//Here we check whether there are any problems with creating a file on the
//server
if pathError != nil{
log.Printf("error creating a file for writing : ", pathError)
return
}
defer out.Close()
//copy content from the file received to the file we created inside the /tmp (server) directory.
_, copyFileError := io.Copy(out, file)
if copyFileError != nil{
log.Printf("error occurred while file copy : ", copyFileError)
}
//Write a message along with a filename to an HTTP response stream
fmt.Fprintf(w, "File uploaded successfully : "+header.Filename)
}

func index(w http.ResponseWriter, r *http.Request){
parsedTemplate, _ := template.ParseFiles("templates/upload-file.html")
parsedTemplate.Execute(w, nil)
}

func main(){
http.HandleFunc("/", index)
http.HandleFunc("/upload", fileHandler)
err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
if err != nil{
log.Fatal("error starting http server : ", err)
return
}
}