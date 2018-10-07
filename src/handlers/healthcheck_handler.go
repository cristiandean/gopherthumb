package handlers

import (
	"fmt"
	"net/http"
)

// HealthcheckHandler: handler healthcheck
func HealthcheckHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, "Working")
	})
}
