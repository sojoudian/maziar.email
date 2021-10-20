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
	log.Fatal(http.ListenAndServeTLS(":443", "localhost.crt", "localhost.key", nil))

}
