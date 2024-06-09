package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

func JSONHandler(jsonFilePath string, fallback http.Handler) (http.HandlerFunc, error) {
	f, err := os.Open(jsonFilePath)
	if err != nil {
		return nil, err
	}

	d := json.NewDecoder(f)
	var urls []pathURL
	d.Decode(&urls)

	mapData := make(map[string]string)

	for _, v := range urls {
		mapData[v.Path] = v.URL
	}

	return MapHandler(mapData, fallback), nil
}
