package main

import (
	"log"
	"net/http"
	"remind-go/handlers"
	"remind-go/routes"
)

func main() {
	StartWeb("9980")
}

func StartWeb(port string) {
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Println("START HTTP Server at " + port)
	go func() {
		handlers.Scheduler()
	}()
	go func() {
		handlers.HandlerErrNotice()
	}()
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("error:" + err.Error())
	}
}
