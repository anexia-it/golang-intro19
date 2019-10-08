package intro19

import (
	"encoding/json"
	"io"
)

type jsonFormat struct{}

func (jsonFormat) Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func (jsonFormat) Encode(w io.Writer, v interface{}) error {
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	return e.Encode(v)
}

func init() {
	ConverterFormats["json"] = jsonFormat{}
}
