package guitars

import (
	"net/http"
	"github.com/vargascardona/neo-ledger/internal/json"
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
	//guitars := struct {
  //  Guitars []string 'json:"guitars'
	//}{}
  guitars := []string{"Gibson", "Fender"}
	json.Write(w, http.StatusOK, guitars)
}
