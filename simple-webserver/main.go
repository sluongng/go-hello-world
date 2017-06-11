package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	//fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("Key:", k)
		fmt.Println("Value:", strings.Join(v, "")) // Convert value to string
	}

	// send data to ResponseWriter, not file
	// assuming used Fprintf for byte-size write buffer
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("[Error]Caught: ", err)
	}
}
