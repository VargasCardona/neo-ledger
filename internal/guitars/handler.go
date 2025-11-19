package guitars

import (
	"net/http"
	"encoding/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListGuitars(w http.ResponseWriter, r *http.Request) {
	guitars := []string{"Gibson", "Fender"}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(guitars)
}
