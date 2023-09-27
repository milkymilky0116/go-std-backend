package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	err := tmp.ExecuteTemplate(os.Stdout, "one.gohtml", "one")
	handleErr(err)
	err = tmp.ExecuteTemplate(os.Stdout, "two.gohtml", "two")
	handleErr(err)
	err = tmp.ExecuteTemplate(os.Stdout, "three.gohtml", "three")
	handleErr(err)
}
func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
