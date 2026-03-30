package handler

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Handler struct {
	store map[string]string
	mu    sync.RWMutex
}

func NewHandler() *Handler {
	return &Handler{
		store: make(map[string]string),
	}
}

func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	type request struct {
		URL string `json:"url"`
	}

	var req request
	_ = json.NewDecoder(r.Body).Decode(&req)

	code := RandString(6)

	h.mu.Lock()
	h.store[code] = req.URL
	h.mu.Unlock()

	resp := map[string]string{"code": code}
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]

	h.mu.RLock()
	url, ok := h.store[code]
	h.mu.RUnlock()

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
