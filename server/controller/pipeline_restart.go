package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/open-devops/pipeline-daemon/server/model"
	"net/http"
)

func RestartPipeline(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	pipelineId := mux.Vars(r)["pipelineId"]
	//capability := mux.Vars(r)["capability"]

	// Get pipeline fundamental info
	pipelineInfo := model.FetchPipelineInfo(pipelineId)

	// Get the pipeline's provisioning & running status
	status := model.FetchPipelineStatus(pipelineInfo)
	response, _ := json.Marshal(status)

	w.Write(response)
}
