package utils

import (
	"net/http"
	"fmt"
)

func ErrorHandle(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	if status == http.StatusNotFound {
		fmt.Fprint(w, "Page not found")
	}
}
