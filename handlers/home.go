package handlers

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter,r *http.Request) {
	fmt.Print("hello")
}
