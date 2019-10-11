package intro19

import (
	"github.com/BurntSushi/toml"
	"io"
	"net/http"
)

type tomlFormat struct{}

func (tomlFormat) Decode(r io.Reader, v interface{}) error {
	_, err := toml.DecodeReader(r, v)
	return err
}

func (tomlFormat) Encode(rw http.ResponseWriter, v interface{}) error {
	e := toml.NewEncoder(rw)
	rw.Header().Set("Content-Type", "application/toml")
	return e.Encode(v)
}

func init() {
	ConverterFormats["toml"] = tomlFormat{}
}
