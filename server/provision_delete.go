package server

import (
	"net/http"
)

type PipelineProvision struct {
}

func DeleteProvision(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
