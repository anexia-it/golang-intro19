package intro19

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

// Format defines a converter format implementation
type Format interface {
	// Decode takes the reader and decodes its contents onto v
	Decode(r io.Reader, v interface{}) error

	// Encode takes v and encodes its contents onto the writer
	Encode(rw http.ResponseWriter, v interface{}) error
}

// ConverterFormats holds all registered formats
var ConverterFormats = make(map[string]Format)

func convertHandler(rw http.ResponseWriter, r *http.Request) {
	// URL needs to contain from and to placeholders
	vars := mux.Vars(r)

	var src, dst Format

	fromFormat := vars["from"]
	toFormat := vars["to"]

	if src = ConverterFormats[fromFormat]; src == nil {
		http.Error(rw, "from format '"+fromFormat+"' not found", http.StatusNotFound)
		return
	}

	if dst = ConverterFormats[toFormat]; dst == nil {
		http.Error(rw, "to format '"+toFormat+"' not found", http.StatusNotFound)
		return
	}

	var data map[string]interface{}

	if err := src.Decode(r.Body, &data); err != nil {
		http.Error(rw, "decoding failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dst.Encode(rw, data); err != nil {
		http.Error(rw, "encoding failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// all done
}
