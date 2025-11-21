package guitars

import (
	"net/http"
	"github.com/vargascardona/neo-ledger/internal/json"
	"github.com/go-chi/chi/v5"
	"log"
	"strconv"
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
	guitars, err := h.service.ListGuitars(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, guitars)
}

func (h *handler) FindGuitarByID(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")

  guitarID, err := strconv.Atoi(id)
    if err != nil {
      http.Error(w, "invalid guitar id", http.StatusBadRequest)
      return
  }

	guitar, err := h.service.FindGuitarByID(r.Context(), guitarID)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, guitar)
}

func (h *handler) CreateGuitar(w http.ResponseWriter, r *http.Request) {
	var tempGuitar Guitar
	if err := json.Read(r,&tempGuitar); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdGuitar, err := h.service.CreateGuitar(r.Context(), tempGuitar)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, createdGuitar)
}

func (h *handler) UpdateGuitar(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")

  guitarID, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
      http.Error(w, "invalid guitar id", http.StatusBadRequest)
      return
  }

	var tempGuitar Guitar
	if err := json.Read(r,&tempGuitar); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tempGuitar.ID = guitarID
	updatedGuitar, err := h.service.UpdateGuitar(r.Context(), tempGuitar)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, updatedGuitar)
}

func (h *handler) DeleteGuitar(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")

  guitarID, err := strconv.Atoi(id)
    if err != nil {
      http.Error(w, "invalid guitar id", http.StatusBadRequest)
      return
  }

	err = h.service.DeleteGuitar(r.Context(), guitarID)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusNoContent, nil)
}

func (h *handler) DisplayName(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")

  guitarID, err := strconv.Atoi(id)
    if err != nil {
      http.Error(w, "invalid guitar id", http.StatusBadRequest)
      return
  }

	guitar, err := h.service.FindGuitarByID(r.Context(), guitarID)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	displayName := h.service.DisplayName(r.Context(), guitar)

	json.Write(w, http.StatusOK, displayName)
}


func (h *handler) IsVintage(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")

  guitarID, err := strconv.Atoi(id)
    if err != nil {
      http.Error(w, "invalid guitar id", http.StatusBadRequest)
      return
  }

	guitar, err := h.service.FindGuitarByID(r.Context(), guitarID)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	isVintage := h.service.IsVintage(r.Context(), guitar)

	json.Write(w, http.StatusOK, isVintage)
}
