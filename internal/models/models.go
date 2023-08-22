package models

type Book struct {
	Id     int
	Title  string
	Author string
}

type Author struct {
	Id         int
	AuthorName string
}
