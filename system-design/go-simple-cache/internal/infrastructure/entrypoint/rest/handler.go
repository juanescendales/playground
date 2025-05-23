package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/juanescendales/playground/system-design/go-simple-cache/internal/infrastructure/adapters"
)

const (
	contentTypeHeader  = "Content-Type"
	appJSONContentType = "application/json"
)

type Handler struct {
	repository *adapters.Repository
}

func NewHandler(repository *adapters.Repository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path == "/key" {
			h.get(w, r)
		} else if r.URL.Path == "/status" {
			h.cacheStatus(w, r)
		}
	case http.MethodPost:
		h.add(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, err := h.repository.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(contentTypeHeader, appJSONContentType)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"value": string(value),
	})
}

func (h *Handler) add(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var createKey CreateKey
	err = json.Unmarshal(body, &createKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.repository.Add(createKey.Key, []byte(createKey.Value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(contentTypeHeader, appJSONContentType)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) cacheStatus(w http.ResponseWriter, _ *http.Request) {
	cacheStatus := h.repository.CacheStatus()
	response := CacheStatusResponse{
		Keys: cacheStatus.OrderedKeys,
		Size: cacheStatus.Size,
	}
	w.Header().Set(contentTypeHeader, appJSONContentType)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
