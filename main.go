package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/curt-labs/IconBoom/handlers"
)

var (
	port = flag.String("port", "8080", "--port=<port> Port to run on")
)

func main() {
	flag.Parse()
	fmt.Print("IconBoom Running. \n")

	http.HandleFunc("/test", handlers.Test)
	http.HandleFunc("/", handlers.Start)

	http.HandleFunc("/insertsuccess", handlers.InsertSuccess)
	http.HandleFunc("/inserterrors", handlers.InsertErrors)

	if port == nil || *port == "" {
		*port = "8080"
	}
	fmt.Print("Port: " + *port)
	http.ListenAndServe(":"+*port, nil)
}
