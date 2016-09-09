package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func DeleteProvision(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	pipelineId := mux.Vars(r)["pipelineId"]
	fmt.Fprintf(w, "Pipeline (ID:"+pipelineId+") deleted!")
}
