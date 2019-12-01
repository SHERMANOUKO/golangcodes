package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

type Post struct{
	ID int
	Title string
	Author string
}

func main(){
	db, err := sql.Open("mysql", "sherman:sherman@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO posts(title, author) VALUES('Post 4','Author 4')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	posts, err := db.Query("SELECT * FROM posts")

	if err != nil {
		panic(err.Error())
	}

	for posts.Next(){
		var post Post
		err := posts.Scan(&post.ID, &post.Title, &post.Author)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(post)
	}
}