package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HealthCheck struct {
}

func (h *HealthCheck) AddRoutes(router *mux.Router) {
	router.HandleFunc("/health", h.handler).Methods(http.MethodGet)
}

func (h *HealthCheck) handler(resp http.ResponseWriter, _ *http.Request) {
	_, _ = resp.Write([]byte(`OK`))
}
