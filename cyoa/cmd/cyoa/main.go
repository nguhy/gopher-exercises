package main

import (
	"flag"
	"fmt"
	"net/http"
	"nguhy/cyoa"
	"os"
)

func main() {

	fileName := flag.String("adventure", "adventure.json", "File that contains the adventure")
	port := flag.Int("port", 8000, "Port to listen to your web server")
	flag.Parse()

	file, err := os.Open(*fileName)

	if err != nil {
		panic(err)
	}

	story, err := cyoa.JSONStory(file)

	if err != nil {
		fmt.Println(fmt.Sprintf("Error parsing file %s", *fileName))
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("Starting Web Server on http://localhost:%d", *port))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), cyoa.NewHandler(story)); err != nil {
		panic(err)
	}
	//fmt.Println(*fileName)
}
