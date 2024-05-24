package handler

import (
	"net/http"

	"github.com/raphaelmb/go-passin/internal/handler/util"
)

type IHealth interface {
	GetHealth(w http.ResponseWriter, r *http.Request)
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"status": "ok",
	}
	util.JSONResponse(w, http.StatusOK, resp)
}
