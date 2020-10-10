// Based on example found at:
// http://stackoverflow.com/questions/17998943/golang-library-package-that-returns-json-string-from-http-request

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	var data struct {
		Items []struct {
			Name            string
			Count           int
			IsRequired      bool
			IsModeratorOnly bool
			HasSynonyms     bool
		}
	}

	r, err := http.Get("https://api.stackexchange.com/2.2/tags?page=1&pagesize=100&order=desc&sort=popular&site=stackoverflow")
	checkErr(err)
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&data)
	checkErr(err)

	for _, item := range data.Items {
		fmt.Printf("%s = %d\n", item.Name, item.Count)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
