package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
	// "html/template"
)

var url string = "https://v2.jokeapi.dev/joke/Programming"

type Joke struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Joke     string `json:"joke"`
	SetUp    string `json:"setup"`
	Delivery string `json:"delivery"`
	ID       int    `json:"id"`
	Safe     bool   `json:"safe"`
}

var templ = `{{range .}}-------------------------
Category: {{.Category}}
Type: {{.Type}}
{{if .Joke}}Joke: {{.Joke}} {{else}}SetUp: {{.SetUp}}
Delivery: {{.Delivery}} {{end}}
ID: {{.ID}}
Safe: {{.Safe}}
{{end}}`

// var templ = `{{range .}}
// <h3>Category: {{.Category}}</h3>
// <h5>Type: {{.Type}}</h5>
// {{if .Joke}}<p>Joke: {{.Joke}}</p> {{else}}<p>Setup: {{.SetUp}}</p>
// <p>Delivery: {{.Delivery}}</p> {{end}}
// <p>ID: {{.ID}},  Safe: {{.Safe}}</p>
// {{end}}`

func main() {
	var jokes []Joke
	for i := 0; i < 10; i++ {
		joke := getJokes()
		jokes = append(jokes, joke)
	}

	report, err := template.New("report").Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	report.Execute(os.Stdout, jokes)
	
}

func getJokes() Joke {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var joke Joke	
	json.NewDecoder(response.Body).Decode(&joke)

	return joke
}
