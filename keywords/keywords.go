// Command keywords prints the name, description, keywords and requirements
// from composer.json.
package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/antonholmquist/jason"
)

func main() {
	filePtr := flag.String("file", "composer.json", "composer.json to read")
	flag.Parse()

	exampleJSON, err := os.ReadFile(filepath.Clean(*filePtr))
	if err != nil {
		panic(err)
	}

	v, err := jason.NewObjectFromBytes(exampleJSON)
	if err != nil {
		log.Fatal(err)
	}

	// Missing optional fields are left empty.
	name, _ := v.GetString("name")
	description, _ := v.GetString("description")

	log.Println("name:", name)
	log.Println("description:", description)

	keywords, _ := v.GetStringArray("keywords")
	for i, keyword := range keywords {
		log.Printf("keyword %d: %s", i, keyword)
	}

	require, err := v.GetObject("require")
	if err != nil {
		return
	}
	for key, value := range require.Map() {
		s, _ := value.String()
		log.Println(key, s)
	}
}
