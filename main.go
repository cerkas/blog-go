package main

import (
	//"blog-go/config"
	"blog-go/models"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	//"database/sql"
	"time"
	//"github.com/globalsign/mgo"
	"github.com/go-martini/martini"
	"html/template"
	"log"
	"net/http"


)
var posts map[string] * models.Post

type Person struct {
	Name string
	Phone string
}
const (
	mongoUrl = "mongodb://myuser:mypass@localhost:27017"
	dbName = "go-blog"
	userCollectionName = ""
)

//var postsCollection  *mgo.Collection
func main() {

	m := martini.Classic()
	posts = make(map[string]*models.Post,0)
	m.Post("/assets/",http.StripPrefix("/assets/",http.FileServer(http.Dir("./assets"))))
	staticOptions := martini.StaticOptions{Prefix:"assets"}
	m.Use(martini.Static("assets",staticOptions))
	m.Get("/", indexHandler)
	m.Get("/news", postHandler)
	m.Get("/write", writeHandler)
	m.Post("/SavePost", savePostHandler)
	log.Println("Listening...")
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	fmt.Println(rows)
	m.Run()
	//fmt.Println(mdb)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "user=user  pass=test dbname=accounts sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	fmt.Println(err)
	fmt.Println(rows)
	//var appConfig config.AppConfig
	fmt.Println(mongoUrl)
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
	t := time.Now()
	date    := t.Format(t.String())
	post    := models.NewPost(id,title,content,date)
	posts[post.Id] = post
	//config.NewSession("localhost:27017")
	/*if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer session.Close()
	userService := root.NewUserService(session.Copy(), dbName, userCollectionName)
	testUsername := "integration_test_user"
	testPassword := "integration_test_password"
	user := root.User{
		Username: testUsername,
		Password: testPassword }

	//Act
	err = userService.Create(&user)*/
	/*if err != nil {
		return nil
	}*/

	/*//Assert
	if err != nil {
		w.Error("Unable to create user: %s", err)
	}
	var results []root.User
	session.GetCollection(dbName,userCollectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Error("Incorrect number of results. Expected `1`, got: `%i`", count)
	}
	if results[0].Username != user.Username {
		t.Error("Incorrect Username. Expected `%s`, Got: `%s`", testUsername, results[0].Username)*/

	/*session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)*/
	fmt.Print(post)
	http.Redirect(w,r,"/",302)
}