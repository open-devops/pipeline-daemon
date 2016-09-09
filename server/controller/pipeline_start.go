package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/open-devops/pipeline-daemon/server/model"
	"net/http"
)

func StartPipeline(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Get pipeline ID from path parameters
	pipelineId := mux.Vars(r)["pipelineId"]
	//capability := mux.Vars(r)["capability"]

	// Get pipeline fundamental info
	pipelineInfo := model.FetchPipelineInfo(pipelineId)

	// Invalid Pipeline ID supplied
	if len(pipelineInfo.PipelineName) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the pipeline's provisioning & running status
	status := model.FetchPipelineStatus(pipelineId)
	response, _ := json.Marshal(status)

	w.Write(response)
}
