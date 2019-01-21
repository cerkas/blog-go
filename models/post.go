package models

type Post struct {
	Id      string
	Title   string
	Сontent string
	Date    string
}

func NewPost( id , title ,content ,date string )* Post  {
	return & Post{ id,title,content,date}
}