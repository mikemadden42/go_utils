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

	v, _ := jason.NewObjectFromBytes(exampleJSON)

	name, _ := v.GetString("name")
	description, _ := v.GetString("name")

	log.Println("name:", name)
	log.Println("description:", description)

	keywords, _ := v.GetStringArray("keywords")
	for i, keyword := range keywords {
		log.Printf("keyword %d: %s", i, keyword)
	}

	require, _ := v.GetObject("require")
	for key, value := range require.Map() {
		s, _ := value.String()
		log.Println(key, s)
	}
}
