package handler

import (
	"encoding/json"
	"net/http"
)

type IHealth interface {
	GetHealth(w http.ResponseWriter, r *http.Request)
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"status": "ok",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
