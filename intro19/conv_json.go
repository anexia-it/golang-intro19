package intro19

import (
	"encoding/json"
	"io"
	"net/http"
)

type jsonFormat struct{}

func (jsonFormat) Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func (jsonFormat) Encode(rw http.ResponseWriter, v interface{}) error {
	e := json.NewEncoder(rw)
	e.SetIndent("", "  ")
	rw.Header().Set("Content-Type", "application/json")
	return e.Encode(v)
}

func init() {
	ConverterFormats["json"] = jsonFormat{}
}
