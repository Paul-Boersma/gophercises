package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
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

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathToUrls, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathToUrls, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func parseYAML(yml []byte) (map[string]string, error) {
	var pathUrls []pathUrl

	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		return nil, err
	}
	pathToUrls := map[string]string{}
	for _, pu := range pathUrls {
		pathToUrls[pu.Path] = pu.Url
	}
	return pathToUrls, nil
}
