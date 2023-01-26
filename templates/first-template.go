package templates
import(

"html/template"
"log"
"net/http"
)

const
(
CONN_HOST = "localhost"
CONN_PORT = "8080"
)
type Person struct{
Id string
Name string
}

func renderTemplate(w http.ResponseWriter, r *http.Request){
person := Person{Id: "1", Name: "Black Eyed Beans"}

//Here we are
// calling ParseFiles of the html/template package, which creates a new template and
// parses the filename we pass as an input, which is first-template.html ,in a templates
// directory. The resulting template will have the name and contents of the input file
parsedTemplate, _ := template.ParseFiles("templates/first-template.html")


// Here we are calling an Execute handler on a
// parsed template, which injects person data into the template, generates an HTML
// output, and writes it onto an HTTP response stream
err := parsedTemplate.Execute(w, person)

if err != nil{
log.Printf("Error occurred while executing the template or writing its output : ", err)
return
}
}

func main(){
http.HandleFunc("/", renderTemplate)
err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
if err != nil{
log.Fatal("error starting http server : ", err)
return
}
}