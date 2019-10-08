package intro19

import (
	"gopkg.in/yaml.v2"
	"io"
	"net/http"
)

type yamlFormat struct{}

func (yamlFormat) Decode(r io.Reader, v interface{}) error {
	return yaml.NewDecoder(r).Decode(v)
}

func (yamlFormat) Encode(rw http.ResponseWriter, v interface{}) error {
	e := yaml.NewEncoder(rw)
	rw.Header().Set("Content-Type", "text/x-yaml")
	return e.Encode(v)
}

func init() {
	ConverterFormats["yaml"] = yamlFormat{}
}
