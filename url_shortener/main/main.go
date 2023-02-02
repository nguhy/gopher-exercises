package main

import (
	"fmt"
	"net/http"
	urlshortener "nguhy/url_shortner"
)

var urlMap = map[string]string{
	"/chat":   "https://chat.openai.com/chat",
	"/gopher": "https://courses.calhoun.io/courses/cor_gophercises",
}

var yamlString = `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	mapHandler := urlshortener.MapHandler(urlMap, mux)

	yamlHandler, err := urlshortener.YamlHandler([]byte(yamlString), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server Starting at port http://localhost:8080")

	http.ListenAndServe(":8080", yamlHandler)
}
