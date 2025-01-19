package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nikubayasi/go-project/cmd/pkg/config"
	"github.com/nikubayasi/go-project/cmd/pkg/handlers"
	"github.com/nikubayasi/go-project/cmd/pkg/render"
)

const portNumber = ":8081"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
