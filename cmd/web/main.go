package main

import (
	"fmt"
	"net/http"

	"github.com/nikubayasi/go-project/cmd/pkg/handlers"
)

const portNumber = ":8081"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
