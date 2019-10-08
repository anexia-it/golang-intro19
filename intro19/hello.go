package intro19

import (
	"encoding/json"
	"net/http"
)

type helloRequest struct {
	Name string `json:"name"`
}

type helloResponse struct {
	Message string
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)

	var r helloRequest
	if err := dec.Decode(&r); err != nil {
		http.Error(rw, "json decoding failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Name == "" {
		r.Name = "world"
	}

	enc := json.NewEncoder(rw)
	enc.Encode(helloResponse{
		Message: "Hello " + r.Name + "!",
	})
}
