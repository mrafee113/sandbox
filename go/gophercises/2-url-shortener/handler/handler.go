package urlshort

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	yamlPkg "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if short, exists := pathsToUrls[r.URL.Path]; exists {
			http.Redirect(w, r, short, http.StatusPermanentRedirect)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	type item struct {
		Path string `yaml:"path"`
		Url  string `yaml:"url"`
	}
	var items []item
	err := yamlPkg.Unmarshal(yml, &items)
	if err != nil {

		return nil, err
	}
	pathsToUrls := make(map[string]string)
	for _, item := range items {
		pathsToUrls[item.Path] = item.Url
	}
	return MapHandler(pathsToUrls, fallback), nil
}

func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	type item struct {
		Path string `json:"path"`
		Url  string `json:"url"`
	}
	var items []item
	err := json.Unmarshal(jsn, &items)
	if err != nil {
		log.Fatalln("Json data could not be parsed.", err)
		return nil, err
	}
	pathsToUrls := make(map[string]string)
	for _, item := range items {
		pathsToUrls[item.Path] = item.Url
	}
	return MapHandler(pathsToUrls, fallback), nil
}

func Sqlite3Handler(filename string, fallback http.Handler) (http.HandlerFunc, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatalln("Could not open database.", err)
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln("Could not ping database.", err)
		return nil, err
	}

	rows, err := db.Query("SELECT path, url FROM data")
	if err != nil {
		log.Fatalln("Could not execute query.", err)
		return nil, err
	}
	defer rows.Close()

	pathsToUrls := make(map[string]string)
	for rows.Next() {
		var path, url string
		err := rows.Scan(&path, &url)
		if err != nil {
			log.Fatalln("Could not scan row.", err)
		}
		pathsToUrls[path] = url
	}

	if err := rows.Err(); err != nil {
		log.Fatalln("There were errors during scanning.", err)
		return nil, err
	}

	return MapHandler(pathsToUrls, fallback), nil
}
