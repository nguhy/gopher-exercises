package url_shortener

import (
	"gopkg.in/yaml.v3"
	"net/http"
)

func MapHandler(urlMap map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if value, ok := urlMap[r.URL.Path]; ok {
			http.Redirect(w, r, value, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func YamlHandler(yamlString []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathURLs []pathURL
	err := parseYAML(yamlString, &pathURLs)
	if err != nil {
		return nil, err
	}

	urlMap := make(map[string]string)
	for _, m := range pathURLs {
		urlMap[m.Path] = m.URL
	}
	return MapHandler(urlMap, fallback), nil
}

func parseYAML(yamlString []byte, pathUrls *[]pathURL) error {
	err := yaml.Unmarshal(yamlString, pathUrls)
	if err != nil {
		return err
	}
	return nil
}
