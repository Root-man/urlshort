package server

import (
	"fmt"
	"net/http"

	"github.com/root-man/urlshort/handlers"
)

func RunJSON(jsonFilePath string) {
	mux := defaultMux()
	handler, err := handlers.JSONHandler(jsonFilePath, mux)
	if err != nil {
		panic(err)
	}

	run(handler)
}

func RunYAML(yamlFilePath string) {
	mux := defaultMux()
	handler, err := handlers.YAMLHandler(yamlFilePath, mux)
	if err != nil {
		panic(err)
	}

	run(handler)
}

func run(handler http.Handler) {
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
