package main

import (
	"fmt"
	"log"
	"net/http"
)
var port = "3001"
var host = "localhost" 
func main() {
	my_router := Router()
	fmt.Println("Starting at %s:%s", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, my_router))
}
