package main

import (
	"log"
	"net/http"
)


func main() {

	// http is a go object that you can register routes like this
	// for example the route for this is "/" which is basically the root
	// the 2nd parameter is what TO DO when someone accesses that route
	// in this example I'm printing something to the console
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		print("I'm printing something on my console...")
	})

	// create the webserver
	// behind the scene, it basically just loops while(true) until someone send a http request to the server
	err := http.ListenAndServe(":8080", nil)

	// If the webserver throws a FATAL error, then we will log it here
	log.Fatal(err)
}


