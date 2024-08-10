package utils

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetPathVar(key string, r *http.Request) string {
	return chi.URLParam(r, key)
}