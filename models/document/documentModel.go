package document

import "time"

//A representation of the document structure
type Document struct{
    ID int
    Owner string
    Body string 
    CreatedOn string
} 

var smapleDocs=[]*Document{
    &Document{ID: 1,
    Owner:"Beans",
    Body: "Identification card",
    CreatedOn: time.Now().String(),},

    &Document{ID: 2,
    Owner:"Kara",
    Body: "Drivers license",
    CreatedOn: time.Now().String(),},
}
