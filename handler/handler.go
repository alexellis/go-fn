package function

import (
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	r.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Handled.. OK"))
}
