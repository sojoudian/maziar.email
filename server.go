package main

import (
	"log"
	"net/http"
)

func main() {

	// handle `/` route
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fs)

	// run server on port "9000"
	log.Fatal(http.ListenAndServe(":80", nil))

}
