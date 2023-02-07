package main

import (
	"flag"
	"fmt"
	"nguhy/cyoa"
	"os"
)

func main() {

	fileName := flag.String("adventure", "adventure.json", "File that contains the adventure")
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

	fmt.Println(fmt.Sprintf("%+v", story))
	fmt.Println(*fileName)
}
