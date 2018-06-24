// http://stackoverflow.com/questions/11805356/text-template-issue-parse-vs-parsefiles

package main

import (
	"os"
	"text/template"
)

type inventory struct {
	Material string
	Count    uint
	Cost     float32
}

func main() {
	sweaters := inventory{"wool", 17, 42.69}
	tmpl, err := template.ParseFiles("template.txt")
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "template.txt", sweaters)
	if err != nil {
		panic(err)
	}
}
