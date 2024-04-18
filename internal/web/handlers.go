package web

import (
	"fmt"
	"net/http"

	"github.com/ethanjmarchand/ezquote/internal/api"
)

type HomeHandler struct {
	Service api.APIService
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s, err := h.Service.Worker.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprint(w, string(s))
}

type ListHandler struct {
	Service api.APIService
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.PathValue("quote")
	s, err := h.Service.Worker.GetQuote(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprint(w, string(s))
}
