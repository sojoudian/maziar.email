package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// handle `/` route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	// run server on port "9000"
	log.Fatal(http.ListenAndServeTLS(":443", "localhost.crt", "localhost.key", nil))

}
