package main

import (
	"fmt"
	"flag"
	"strings"
	"io/ioutil"
	"html/template"
	"os"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type Page struct {
	Content string
}

func main() {
	dir := flag.String("dir", " ", "directory path flag")
	flag.Parse()


	files, err := ioutil.ReadDir(*dir)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := file.Name()
		dotIndex := strings.Index(filePath, ".")

		var fileName string
		var fileExtension string

		if dotIndex == -1 {
			continue
		} else {
			fileName = filePath[:dotIndex]
			fileExtension = filePath[dotIndex:]
		}

		if fileExtension == ".txt" || fileExtension == ".md" {
			fmt.Println(fileName + fileExtension)

			if fileExtension == "txt" {
				data, err := ioutil.ReadFile("./" + fileName + ".txt")
		
				if err != nil {
					panic(err)
				}
			
				page := Page{Content: string(data)}
				t, _ := template.ParseFiles("template.tmpl")
				newFile, _ := os.Create(fileName + ".html")
				err = t.Execute(newFile, page)
	
				if err != nil {
					panic(err)
				}
			}
		} else if fileExtension == ".md" {
			data, err := ioutil.ReadFile("./" + fileName + ".md")

			if err != nil {
				panic(err)
			}

			extensions := parser.CommonExtensions | parser.AutoHeadingIDs
			parser := parser.NewWithExtensions(extensions)

			md := []byte(data)
			output := markdown.ToHTML(md, parser, nil)

			fileContents := template.HTML(string(output))


			markdownPage := Page{Content: string(fileContents)}
			t, _ := template.ParseFiles("template.tmpl")
			markdownFile, _ := os.Create(fileName + ".html")
			err = t.Execute(markdownFile, markdownPage)

			if err != nil {
				panic(err)
			}

		}
	}
}