package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/open-devops/pipeline-daemon/server/types"
	"net/http"
)

func GetProvision(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	pipelineId := mux.Vars(r)["pipelineId"]

	provision := &types.PipelineProvision{
		PipelineId:   pipelineId,
		Capabilities: nil,
	}

	response, _ := json.Marshal(provision)
	w.Write(response)
}
