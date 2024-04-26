package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ethanjmarchand/ezquote/internal/quoting"
)

type HomeHandler struct {
	Service quoting.QuotingService
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s, err := h.Service.Worker.GetAll()
	if err != nil {
		// Could potentially run a switch statement depending on the error.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	bs, err := json.Marshal(s)
	if err != nil {
		fmt.Println("error marshaling the json")
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}

type ListHandler struct {
	Service quoting.QuotingService
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.PathValue("quote")
	if q == "" {
		fmt.Println("user provided an invalid quote string")
		http.Error(w, "Must provide a valid quote ID", http.StatusBadRequest)
		return
	}
	s, err := h.Service.Worker.GetQuote(q)
	if err != nil {
		if errors.Is(err, quoting.ErrInvalidCoinID) {
			http.Error(w, "Invalid Token ID", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bs, err := json.Marshal(s)
	if err != nil {
		fmt.Println("error marshaling the json")
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}
