package main

import (
	"encoding/json"
	"excercises/cyoa/types"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var t *template.Template

type handle struct {
	book *types.Book
}

func (h handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for key := range *h.book {
		if strings.Split(r.URL.Path, "/")[1] == key {
			t.Execute(w, (*h.book)[key])
			return
		}
	}
	t.Execute(w, (*h.book)["intro"])

}

func init() {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<style>
			html,body{
				display:flex;
				justify-content:center;
				align-items:center;
				height:100%;
			}
			p,{
				margin:5% 0
			}
			h2{
				text-align:center;
			}
			
			
			.story-container {
				margin: 0 15%;
				padding: 20% 5%;
				box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
				border-radius:5px;
				background:aliceblue;
			}
		</style>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
	<div class="story-container">
	<h2>{{.Title}}</h2>
	{{range .Story}}<p> {{.}}</p>{{end}}
	{{range .Options}}<a href="/{{.Arc}}"> {{.Text}}</a><br />{{end}}	
	</div>
	</body>	
</html>`
	t = template.Must(template.New("webpage").Parse(tpl))
}

func main() {
	//open json file
	//read json file
	//unmarshal json file
	book := &types.Book{}
	parseJSON(book)
	log.Fatal(http.ListenAndServe(":3000", handle{book}))

}

func parseJSON(book *types.Book) {
	fd, err := os.Open("../../gophers.json")
	if err != nil {
		log.Fatalf("Unable to open json file: %v", err)
	}
	defer fd.Close()
	jsondata, err := ioutil.ReadAll(fd)
	if err != nil {
		log.Fatalf("Unable to read json file: %v", err)
	}
	//
	if err := json.Unmarshal(jsondata, book); err != nil {
		log.Fatalf("Unable to parse json file: %v", err)
	}
}
