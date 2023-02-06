package main

import (
	"encoding/json"
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

	dec := json.NewDecoder(file)

	var story cyoa.Story

	err = dec.Decode(&story)

	if err != nil {
		fmt.Println(fmt.Sprintf("Error parsing file %s", fileName))
		return
	}

	fmt.Println(fmt.Sprintf("%+v", story))
	fmt.Println(*fileName)
}
