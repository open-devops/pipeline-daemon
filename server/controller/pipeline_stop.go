package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/open-devops/pipeline-daemon/server/model"
	"net/http"
)

func StopPipeline(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	pipelineId := mux.Vars(r)["pipelineId"]
	capability := mux.Vars(r)["capability"]

	// Get pipeline fundamental info
	pipelineInfo := model.FetchPipelineInfo(pipelineId)

	// Stop the pipeline's provisioning & Fetch running status
	if status, err := model.StartPipeline(pipelineInfo, capability); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		response, _ := json.Marshal(status)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
