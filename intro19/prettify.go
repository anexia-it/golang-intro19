package intro19

import (
	"github.com/gorilla/mux"
	"net/http"
)

func prettifyHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vars["from"] = vars["fmt"]
	vars["to"] = vars["fmt"]
	mux.SetURLVars(r, vars)
	convertHandler(rw, r)
}
