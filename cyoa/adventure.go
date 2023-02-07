package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"
)

func init() {
	temp = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

var defaultHandlerTemplate = `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>Title</title>
			</head>
			<body>
				<h1>{{.Title}}</h1>
				{{range .Paragraphs}}
					<p>{{.}}</p>
				{{end}}
				
				<ul>
					{{range .Options}}
						<li>
							<a href="{{.Chapter}}">{{.Text}}</a>
						</li>
					{{end}}
				</ul>
			</body>
		</html>
	`

var temp *template.Template

type handler struct {
	story Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := temp.Execute(w, h.story["intro"])
	if err != nil {
		panic(err)
	}
}

func NewHandler(s Story) http.Handler {
	return handler{
		story: s,
	}
}

type Story map[string]Chapter

func JSONStory(r io.Reader) (Story, error) {

	dec := json.NewDecoder(r)

	var story Story

	if err := dec.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
