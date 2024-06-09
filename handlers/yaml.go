package handlers

import (
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

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
func YAMLHandler(yamlFilePath string, fallback http.Handler) (http.HandlerFunc, error) {
	var pathURLs []pathURL

	f, err := os.Open(yamlFilePath)
	if err != nil {
		return nil, err
	}

	d := yaml.NewDecoder(f)

	if err := d.Decode(&pathURLs); err != nil {
		return nil, err
	}

	mapHandlerData := make(map[string]string)
	for _, v := range pathURLs {
		mapHandlerData[v.Path] = v.URL
	}

	return MapHandler(mapHandlerData, fallback), nil
}
