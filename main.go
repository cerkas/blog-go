package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"awesomeProject1/models"
)
var posts map[string] * models.Post
func main() {
	posts = make(map[string]*models.Post,0)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/news", postHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/SavePost", savePostHandler)
	http.Handle("/assets/",http.StripPrefix("/assets/",http.FileServer(http.Dir("./assets"))))
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html","templates/header.html","templates/footer.html")
	if err !=nil {
		fmt.Printf( err.Error())
	}
	t.ExecuteTemplate(w,"index",nil)
}
func postHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/news.html","templates/header.html","templates/footer.html")
	if err !=nil {
		fmt.Printf( err.Error())
    }
	t.ExecuteTemplate(w,"news",posts)
	}
func writeHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/write.html","templates/header.html","templates/footer.html")
	if err !=nil {
		fmt.Printf( err.Error())
	}
	t.ExecuteTemplate(w,"write",nil)
}
func savePostHandler(w http.ResponseWriter, r *http.Request)  {
	id      := r.FormValue("id")
	title   := r.FormValue("title")
	content := r.FormValue("content")
	post    := models.NewPost(id,title,content)
	posts[post.Id] = post
	fmt.Print(post)
	http.Redirect(w,r,"/",302)
}