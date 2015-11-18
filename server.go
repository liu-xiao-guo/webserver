package main

import (
	"github.com/go-martini/martini"
	"log"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Static("public"))

	m.Get("/", func() string {
		return "Hello world!"
	})

	if err := http.ListenAndServe(":8001", m); err != nil {
		log.Fatal(err)
	}

	m.Run()
}
