package main

import (
	"awesomeProject1/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/go-martini/martini"
)
var posts map[string] * models.Post
func main() {
	m := martini.Classic()
	fmt.Println(m)
	posts = make(map[string]*models.Post,0)
	m.Post("/assets/",http.StripPrefix("/assets/",http.FileServer(http.Dir("./assets"))))
	staticOptions := martini.StaticOptions{Prefix:"assets"}
	m.Use(martini.Static("assets",staticOptions))
	m.Get("/", indexHandler)
	m.Get("/news", postHandler)
	m.Get("/write", writeHandler)
	m.Post("/SavePost", savePostHandler)

	log.Println("Listening...")
	//http.ListenAndServe(":3000", nil)//Default port l
	m.Run()
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