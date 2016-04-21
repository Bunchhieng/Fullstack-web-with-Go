package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "Bunchhieng"}
	f2 := Friend{Fname: "Siharany"}
	t := template.New("fieldname example")
	// {{.}} point to current object
	t, _ = t.Parse(`hello {{.UserName}}!
	{{range .Emails}}
		an email {{.}}
	{{end}}
	{{with .Friends}}
	{{range .}}
		my friend name is {{.Fname}}
	{{end}}
	{{end}}
	`)
	// {{if `anything`}} {{end}} if-else only support boolean
	// {{. | html}} escape to html using pipe |
	p := Person{UserName: "Bunchhieng", Emails: []string{"Bunchhieng@gmail.com"}, Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
