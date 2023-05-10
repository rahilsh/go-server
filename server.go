package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Fprint(w, `{
    			"total": 0,
    			"count": 0,
    			"books": []
			}`)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `Hello world!`)
	})
	http.HandleFunc("/books", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
