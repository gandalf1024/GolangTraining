package main

import (
	"html/template"
	"log"
	"os"
	"fmt"
)

type Page struct {
	Title string
	Body  string
}

func main() {

	tpl, err := template.ParseFiles("tpl.gohtml", "tpl2.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, Page{
		Title: "Which file?",
		Body:  "hello page 1",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n***************")

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", Page{
		Title: "specifying tpl.gohtml",
		Body:  "hello page 1",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n***************")

	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", Page{
		Title: "specifying tpl2.gohtml",
		Body:  "hello page 2",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
