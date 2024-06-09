package handlers

import "net/http"

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		if to, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, to, http.StatusMovedPermanently)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
