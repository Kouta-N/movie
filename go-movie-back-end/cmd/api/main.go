package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// set application config
	var app application

	// read from command line

	// connect to the database
	app.Domain = "example.com"

	log.Println("Starting application in port", port)

	// 第1引数がパス、第2引数がハンドラ
	http.HandleFunc("/", Hello)

	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
