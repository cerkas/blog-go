package main

import (
	"fmt"
	//"fmt"
	"html/template"
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/", postHandler)
	http.Handle("/assets/",http.StripPrefix("/assets/",http.FileServer(http.Dir("./assets"))))
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "<h1>Hello Wougjugrld</h1>")
	t, err := template.ParseFiles("templates/index.html","templates/header.html","templates/footer.html")
	if err !=nil {
		fmt.Printf( err.Error())
    }
	t.ExecuteTemplate(w,"index",nil)
	}
