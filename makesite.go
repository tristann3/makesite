package main

import (
	// "fmt"
	"html/template"
	"os"
)

type Post struct {
	Content string
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	dat, err := os.ReadFile("./first-post.txt")
    check(err)

	post := Post{Content: string(dat)}
	parsedTemplate, _ := template.ParseFiles("template.tmpl")
	newFile, _ := os.Create("first-post.html")
	err = parsedTemplate.Execute(newFile, post)
	check(err)
}