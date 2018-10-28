package models

type Post struct {
	Id      string
	Title   string
	Сontent string
}

func NewPost( id , title ,content string )* Post  {
	return & Post{ id,title,content}
}