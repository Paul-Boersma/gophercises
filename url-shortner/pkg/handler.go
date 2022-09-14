package urlshort

import (
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	// MapHandler handles all routes in a similar fashion and is therefore best used as
	// encapsulation function. Is this correct?
	// Yes and no, but technically no. MapHandler redirects requests, meaning another request is
	// transmitted to a url that is actually able to handle the provided path.
	// In this case the MapHandler redirects to 3rd party websites, so no
	// specific handler implementation is required on our server.
	// The MapHandler also takes a fallback ServeMux, which currently implements the route '/',
	// the handler for this route has been implemented in main.go
	return func(rw http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if dest, ok := pathsToUrls[url]; ok {
			http.Redirect(rw, r, dest, http.StatusFound)
		}
		fallback.ServeHTTP(rw, r)
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
	// TODO: Implement this...
	return nil, nil
}
