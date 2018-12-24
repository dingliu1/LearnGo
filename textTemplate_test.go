package main

import (
	"fmt"
	xmlT "html/template"
	"log"
	"os"
	"strings"
	"testing"
	textT "text/template"
)

func TestTextTemplate1(t *testing.T) {

	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, err := textT.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	g, err := textT.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = g.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
	fmt.Println()
	k, err := xmlT.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = k.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")

}

func TestTextTemplate2(t *testing.T) {
	// Define a template.
	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	d := textT.Must(textT.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := d.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}

func TestTextTemplate3(t *testing.T) {
	const (
		master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list"}} {{join . ", "}}{{end}} `
	)
	var (
		funcs     = textT.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	masterTmpl, err := textT.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}
	overlayTmpl, err := textT.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
}

func TestHtmlTemplate(t *testing.T) {

	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	d, err := xmlT.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = d.Execute(os.Stdout, data)
	check(err)

	noItems := struct {
		Title string
		Items []string
	}{
		Title: "My another page",
		Items: []string{},
	}

	err = d.Execute(os.Stdout, noItems)
	check(err)

	fmt.Println(EquableTriangle(5, 12, 13))

}

func EquableTriangle(a, b, c int) bool {

	return c*c == a*a+b*b || a*a == c*c+b*b || b*b == c*c+a*a

}
