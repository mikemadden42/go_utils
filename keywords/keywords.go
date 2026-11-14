package main

import (
	"log"
	"os"

	"github.com/antonholmquist/jason"
)

func main() {

	exampleJSON, err := os.ReadFile("composer.json")
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
