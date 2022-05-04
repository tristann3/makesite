package main

import (
	"fmt"
	"flag"
	"strings"
	"io/ioutil"
	"html/template"
	"os"
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

		if fileExtension == ".txt" {
			fmt.Println(fileName + fileExtension)

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
	}
}