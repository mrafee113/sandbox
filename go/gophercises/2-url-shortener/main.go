package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	urlshort "url-shortener/handler"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	var filename string
	flag.StringVar(&filename, "filename", "paths-to-urls.yaml", "input filename")
	flag.Parse()

	var handler http.HandlerFunc
	extension := filepath.Ext(filename)
	if strings.ToLower(extension) == ".json" {
		data, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalln("FATAL. Failed to open/read file.", filename, err)
			panic(err)
		}
		jsonHandler, err := urlshort.JSONHandler([]byte(data), mapHandler)
		if err != nil {
			panic(err)
		}
		handler = jsonHandler
	} else if strings.ToLower(extension) == ".sqlite3" {
		sqlHandler, err := urlshort.Sqlite3Handler(filename, mapHandler)
		if err != nil {
			panic(err)
		}
		handler = sqlHandler
	} else {
		data, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalln("FATAL. Failed to open/read file.", filename, err)
			panic(err)
		}
		yamlHandler, err := urlshort.YAMLHandler([]byte(data), mapHandler)
		if err != nil {
			panic(err)
		}
		handler = yamlHandler
	}
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
