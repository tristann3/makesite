package main

import (
	"fmt"
	"html/template"
	"os"
	"flag"
	"strings"
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
	file := flag.String("file", " ", "parse file path")
	flag.Parse()
	fileName := strings.Split(*file, ".")[0]

	dat, err := os.ReadFile("./" + fileName + ".txt")
    check(err)

	post := Post{Content: string(dat)}
	parsedTemplate, _ := template.ParseFiles("template.tmpl")
	newFile, _ := os.Create(fileName + ".html")
	err = parsedTemplate.Execute(newFile, post)
	check(err)

	fmt.Println("Done")
}