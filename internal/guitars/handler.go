package guitars

import (
	"net/http"
	"github.com/vargascardona/neo-ledger/internal/json"
	"log"
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
	err := h.service.ListGuitars(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	guitars := struct {
    Guitars []string `json:"guitars"`
	}{}

	json.Write(w, http.StatusOK, guitars)
}
