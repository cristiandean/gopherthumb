package handlers

import "github.com/gorilla/mux"

// HTTPHandler Handle http router
func HTTPHandler(router *mux.Router) {
	router.Handle("/", ImageHandler())
	router.PathPrefix("/unsafe/").Handler(ImageHandler())
	router.Handle("/healthcheck", HealthcheckHandler())
}
