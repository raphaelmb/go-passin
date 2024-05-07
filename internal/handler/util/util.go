package util

import (
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload any) {}

func RespondWithError(w http.ResponseWriter, code int, payload any) {}
