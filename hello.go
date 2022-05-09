package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(res http.ResponseWriter, req *http.Request) {
	var message = fmt.Sprintf("Hello, %s!", req.URL.Path[1:])
	log.Print(message)
	fmt.Fprintln(res, message)
}

func main() {
	var PORT = 80
	http.HandleFunc("/", HelloServer)
	log.Println("Hello world server starting on", fmt.Sprintf("http://localhost:%d", PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
